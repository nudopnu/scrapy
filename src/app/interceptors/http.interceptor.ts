import { HttpInterceptorFn, HttpStatusCode } from "@angular/common/http";
import { inject } from "@angular/core";
import { catchError, of, switchMap, throwError } from "rxjs";
import { AuthService } from "../services/auth.service";

export const httpInterceptor: HttpInterceptorFn = (req, next) => {
  const auth = inject(AuthService);
  const accessToken = auth.accessToken;
  const refreshToken = auth.getRefreshToken();
  console.log("[INTERCEPTOR]", req);
  if (accessToken) {
    console.log(auth.parseJWT(accessToken));
    req = req.clone({
      headers: req.headers.set("Authorization", `Bearer ${accessToken}`),
    });
  }
  return next(req)
    .pipe(
      catchError((error) => {
        console.log(error);
        if (error.status === 0 ){
          return throwError(() => ({
            error: new Error("The server seems to be offline"),
          }));
        }
        if (error.status != HttpStatusCode.Unauthorized || !refreshToken) {
          return throwError(() => error);
        }
        return auth.refreshToken(refreshToken).pipe(
          catchError((_) => {
            auth.logout();
            return throwError(() => ({
              error: new Error("invalid refresh token"),
            }));
          }),
          switchMap((res) => {
            req = req.clone({
              setHeaders: { Authorization: `Bearer ${res.access_token}` },
            });
            return next(req);
          }),
        );
      }),
    );
};
