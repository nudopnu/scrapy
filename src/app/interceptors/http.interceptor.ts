import {
  HttpErrorResponse,
  HttpInterceptorFn,
  HttpResponse,
  HttpStatusCode,
} from "@angular/common/http";
import { inject } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, tap, throwError } from "rxjs";
import { AuthService } from "../services/auth.service";
import { LoginResponse } from "../models/responses";

export const httpInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  const auth = inject(AuthService);
  const accessToken = auth.accessToken;
  if (accessToken) {
    req.headers.set("Authorization", `Bearer ${accessToken}`);
  }
  return next(req).pipe(
    tap((event) => {
      if (event instanceof HttpResponse && event.body) {
        const body = event.body as LoginResponse;
        // Assuming the auth token and user info are in the response body
        if (body) {
          auth.setAccessToken(body.access_token);
          auth.setUserInfo(body.username); // Adjust according to your response structure
        }
      }
    }),
    catchError((error) => {
      console.log(error);
      if (error instanceof HttpErrorResponse) {
        if (error.status == HttpStatusCode.Unauthorized) {
          const oldPath = `/${(router as any).location._basePath}`;
          router.navigate(["login"], { queryParams: { redirectTo: oldPath } });
        }
      }
      if (error.error.message) {
        return throwError(() => error.error);
      } else {
        return throwError(() => ({ message: error.statusText }));
      }
    }),
  );
};
