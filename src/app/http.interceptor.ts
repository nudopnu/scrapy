import { HttpErrorResponse, HttpInterceptorFn, HttpStatusCode } from "@angular/common/http";
import { inject } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, of } from "rxjs";

export const httpInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  return next(req).pipe(catchError((error) => {
    if (error instanceof HttpErrorResponse) {
      if (error.status == HttpStatusCode.Unauthorized)       {
        router.navigate(['login']);
      }
    }
    return of(error);
  }));
};
