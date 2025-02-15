import {
	ApplicationConfig,
	ErrorHandler,
	LOCALE_ID,
	provideZoneChangeDetection,
} from '@angular/core';
import {
	provideRouter,
	withComponentInputBinding,
	withRouterConfig,
} from '@angular/router';

import {
	HTTP_INTERCEPTORS,
	provideHttpClient,
	withInterceptorsFromDi,
} from '@angular/common/http';
import { routes } from './app.routes';
import { AuthInterceptor } from './interceptors/auth.interceptor';

import { registerLocaleData } from '@angular/common';
import localeDe from '@angular/common/locales/de';
import { ErrorService } from './services/error.service';

registerLocaleData(localeDe);

export const appConfig: ApplicationConfig = {
	providers: [
		provideZoneChangeDetection({ eventCoalescing: true }),
		provideRouter(
			routes,
			withComponentInputBinding(),
			withRouterConfig({ onSameUrlNavigation: 'reload' }),
		),
		provideHttpClient(withInterceptorsFromDi()),
		{ provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
		{ provide: LOCALE_ID, useValue: 'de-DE' },
		{ provide: ErrorHandler, useClass: ErrorService },
	],
};
