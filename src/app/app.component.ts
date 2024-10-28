import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { initFlowbite } from 'flowbite';
import { NavigationComponent } from './components/navigation/navigation.component';

@Component({
  selector: 'fs-root',
  standalone: true,
  imports: [RouterOutlet, NavigationComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit{
  title = 'flowbite-sample';
  
  ngOnInit(): void {
    initFlowbite();
  }
}
