import { HttpInterceptorFn } from "@angular/common/http";
import { inject } from "@angular/core";
import { AuthService } from "../services/auth.service";

export const httpInterceptor: HttpInterceptorFn = (req, next) => {
  const auth = inject(AuthService);
  const accessToken = auth.accessToken;
  console.log("[INTERCEPTOR]", req);
  if (accessToken) {
    console.log(auth.parseJWT(accessToken));
    req = req.clone({
      headers: req.headers.set("Authorization", `Bearer ${accessToken}`),
    });
  }
  return next(req);
  // .pipe(
  //   catchError((error) => {
  //     console.log(error);
  //     if (error instanceof HttpErrorResponse) {
  //       if (error.status == HttpStatusCode.Unauthorized) {
  //         if (refreshToken) {
  //           return auth.refreshToken(refreshToken).pipe(
  //             switchMap((res) => {
  //               req = req.clone({
  //                 setHeaders: { Authorization: `Bearer ${res.access_token}` },
  //               });
  //               return next(req);
  //             }),
  //           );
  //         }
  //         const oldPath = `/${(router as any).location._basePath}`;
  //         router.navigate(["login"], { queryParams: { redirectTo: oldPath } });
  //       }
  //     }
  //     if (error.error.message) {
  //       return throwError(() => error.error);
  //     } else {
  //       return throwError(() => ({ message: error.statusText }));
  //     }
  //   }),
  // );
};
