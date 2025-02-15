import { Component, inject, input, output, ViewRef } from '@angular/core';

@Component({
	selector: 'fs-alert',
	imports: [],
	templateUrl: './alert.component.html',
	styleUrl: './alert.component.css',
})
export class AlertComponent {
	message = input.required<string>();
	destroy = output();
}
