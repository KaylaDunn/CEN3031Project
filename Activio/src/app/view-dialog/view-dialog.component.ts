import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-view-dialog',
  templateUrl: './view-dialog.component.html',
  styleUrls: ['./view-dialog.component.css']
})
export class ViewDialogComponent {

  post: any;
  constructor(private http: HttpClient, @Inject(MAT_DIALOG_DATA) public data: { id: number }) { }

  ngOnInit(): void {
    // get post with that id
    this.http.get<any>(`http://localhost:3000/api/post/${this.data.id}`)
    .subscribe(response => {

      this.post = response.post;
      
    }, error => {
      console.error(error);
    });
  }
}
