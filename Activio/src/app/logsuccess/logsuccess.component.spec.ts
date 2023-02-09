import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogsuccessComponent } from './logsuccess.component';

describe('LogsuccessComponent', () => {
  let component: LogsuccessComponent;
  let fixture: ComponentFixture<LogsuccessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LogsuccessComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LogsuccessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
