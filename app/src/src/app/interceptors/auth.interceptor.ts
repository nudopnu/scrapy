import {
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
  HttpStatusCode,
} from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { catchError, Observable, switchMap, throwError } from "rxjs";
import { AuthService } from "../services/auth.service";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  auth = inject(AuthService);
  isRefreshing = false;

  intercept(
    req: HttpRequest<any>,
    next: HttpHandler,
  ): Observable<HttpEvent<any>> {
    const accessToken = this.auth.accessToken;
    const refreshToken = this.auth.getRefreshToken();
    if (
      accessToken &&
      Math.floor(Date.now() / 1000) < this.auth.parseJWT(accessToken).exp
    ) {
      req = req.clone({
        headers: req.headers.set("Authorization", `Bearer ${accessToken}`),
      });
    }
    return next.handle(req).pipe(catchError((err) => {
      if (
        err.status === HttpStatusCode.Unauthorized && !this.isRefreshing &&
        refreshToken
      ) {
        this.isRefreshing = true;
        return this.auth.refreshToken(refreshToken)
          .pipe(
            switchMap((res) => {
              this.isRefreshing = false;
              req = req.clone({
                setHeaders: { Authorization: `Bearer ${res.access_token}` },
              });
              return next.handle(req);
            }),
          );
      }
      return throwError(() => err);
    }));
  }
}