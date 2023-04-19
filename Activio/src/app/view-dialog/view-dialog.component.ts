import { Component, Inject } from '@angular/core';
import { MatDialog, MatDialogConfig, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';
import { CommentDialogComponent } from '../comment-dialog/comment-dialog.component';
import { AddImageComponent } from '../add-image/add-image.component';

@Component({
  selector: 'app-view-dialog',
  templateUrl: './view-dialog.component.html',
  styleUrls: ['./view-dialog.component.css']
})
export class ViewDialogComponent {

  post: any;
  userID: number;
  constructor(private http: HttpClient, @Inject(MAT_DIALOG_DATA) public data: { id: number }, private dialog: MatDialog) { 
    this.userID = -1;
  }

  ngOnInit(): void {
    // get post with that id
    this.http.get<any>(`http://localhost:3000/api/post/${this.data.id}`)
    .subscribe(response => {

      this.post = response.post;
      
    }, error => {
      console.error(error);
    });

    // get current users info
    this.http.get<any>("http://localhost:3000/api/auth/me", {
      withCredentials: true
    })
    .subscribe(response => {
      this.userID = response.id;
    }, error => {
      console.error(error);
    });
  }

  onLike() {

    this.http.put(`http://localhost:3000/api/auth/likepost/${this.post.id}`, null, {
      withCredentials: true
    })
    .subscribe(response => {
      console.log(response);
      window.location.reload();
    }, error => {
      console.error(error);
    }
    )
  }

  onComment() {
    const dialogRef: MatDialogRef<CommentDialogComponent> = this.dialog.open(CommentDialogComponent);

    dialogRef.afterClosed().subscribe((comment) => {
      if (comment) {
        console.log("Comment: ", comment);
        this.http.put(`http://localhost:3000/api/auth/comment/${this.post.id}`,
        {
          comment
        },
        {
          withCredentials: true
        })
        .subscribe(response => {
          console.log(response);
          window.location.reload();
        }, (error) => {
          console.log(error);
        });
      }
    })
  }

  addImage() {

    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = { id: this.post.id };
    const dialogRef: MatDialogRef<AddImageComponent> = this.dialog.open(AddImageComponent, dialogConfig);

    dialogRef.afterClosed().subscribe(result => {
      // reload parent component
      location.reload();
    });
    
  }
}
