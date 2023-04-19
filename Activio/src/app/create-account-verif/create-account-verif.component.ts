import { Component } from '@angular/core';

@Component({
  selector: 'app-create-account-verif',
  templateUrl: './create-account-verif.component.html',
  styleUrls: ['./create-account-verif.component.css']
})
export class CreateAccountVerifComponent {
  constructor() { }

  ngOnInit(): void {
    setTimeout(() => {
      window.location.href = 'http://localhost:4200/login';
    }, 5000); // 5000 milliseconds = 5 seconds
  }
}
