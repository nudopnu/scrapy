import { Component, inject, OnChanges, SimpleChanges } from '@angular/core';
import { initFlowbite } from 'flowbite';
import { AuthService } from '../../services/auth.service';
import { BottomNavigationComponent } from './bottom-navigation/bottom-navigation.component';
import { TopNavigationComponent } from './top-navigation/top-navigation.component';

@Component({
	selector: 'fs-navigation',
	imports: [TopNavigationComponent, BottomNavigationComponent],
	templateUrl: './navigation.component.html',
})
export class NavigationComponent implements OnChanges {
	auth = inject(AuthService);
	links = [{ path: '/', label: 'Home' }];

	ngOnChanges(changes: SimpleChanges): void {
		initFlowbite();
	}
}
