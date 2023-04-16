
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-logout',
  template: `
    <h1>You have been logged out.</h1>
    <p>Redirecting to the main page in 5 seconds...</p>
  `,
})
export class LogoutComponent implements OnInit {
  constructor() { }

  ngOnInit(): void {
    setTimeout(() => {
      window.location.href = 'header.html';
    }, 5000); // 5000 milliseconds = 5 seconds
  }
}




