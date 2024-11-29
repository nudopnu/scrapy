import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdsRouteComponent } from './ads-route.component';

describe('AdsRouteComponent', () => {
  let component: AdsRouteComponent;
  let fixture: ComponentFixture<AdsRouteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AdsRouteComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AdsRouteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
