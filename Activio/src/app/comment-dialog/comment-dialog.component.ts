import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-comment-dialog',
  templateUrl: './comment-dialog.component.html',
  styleUrls: ['./comment-dialog.component.css']
})

export class CommentDialogComponent {
  comment: string;
  constructor(private dialogRef: MatDialogRef<CommentDialogComponent>, @Inject(MAT_DIALOG_DATA) public data: any) {
    this.comment = '';
  }

  closeDialog() {
    this.dialogRef.close(this.comment);
  }
}
