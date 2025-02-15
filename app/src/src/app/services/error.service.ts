import {
	ApplicationRef,
	createComponent,
	EnvironmentInjector,
	ErrorHandler,
	inject,
	Injectable,
} from '@angular/core';
import { AlertComponent } from '../components/alert/alert.component';

@Injectable({
	providedIn: 'root',
})
export class ErrorService implements ErrorHandler {
	injector = inject(EnvironmentInjector);

	async handleError(error: any) {
		const appRef = this.injector.get(ApplicationRef);
		const componentRef = createComponent(AlertComponent, {
			environmentInjector: this.injector,
		});
		appRef.attachView(componentRef.hostView);
		const xhr = error.target as any;

		let message =
			error instanceof ProgressEvent
				? `Request failed with error type: ${error.type}. URL: ${xhr['__zone_symbol__xhrURL'] || 'Unknown URL'}.`
				: 'An unexpected error occurred';
		if (error.message) {
			message = error.message;
		}
		componentRef.setInput('message', message);
		componentRef.instance.destroy.subscribe(() => componentRef.destroy());
		document.body.appendChild((componentRef.hostView as any).rootNodes[0]);
		console.error(error);
	}
}
