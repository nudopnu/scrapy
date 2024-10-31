import {
  HttpErrorResponse,
  HttpInterceptorFn,
  HttpStatusCode,
} from "@angular/common/http";
import { inject } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, retry, switchMap, throwError } from "rxjs";
import { AuthService } from "../services/auth.service";

export const httpInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  const auth = inject(AuthService);
  const accessToken = auth.accessToken;
  const refreshToken = auth.getRefreshToken();
  console.log(req);
  if (accessToken) {
    req = req.clone({
      headers: req.headers.set("Authorization", `Bearer ${accessToken}`),
    });
  }
  return next(req).pipe(
    catchError((error) => {
      console.log(error);
      if (error instanceof HttpErrorResponse) {
        if (error.status == HttpStatusCode.Unauthorized) {
          if (refreshToken) {
            return auth.refreshToken(refreshToken).pipe(
              switchMap((res) => {
                req = req.clone({
                  setHeaders: { Authorization: `Bearer ${res.access_token}` },
                });
                return next(req);
              }),
            );
          }
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
