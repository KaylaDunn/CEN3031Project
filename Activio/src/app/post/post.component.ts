import { HttpClient, HttpEventType } from '@angular/common/http';
import { Component, Inject } from '@angular/core';
import * as _ from 'cypress/types/lodash';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent {
  selectedFile: File | null = null;

  constructor(private http: HttpClient) { }

  onFileSelected(event: any) {
    this.selectedFile = <File>event?.target.files[0];
  }

  onUpload() {
    if (this.selectedFile) {
      const fd = new FormData()
      fd.append('image', this.selectedFile, this.selectedFile.name);
      this.http.post('aa', fd, { // add url to backend function that accepts foreign data
        reportProgress: true,
        observe: 'events'
      }) 
      .subscribe(event => {
          if (event.type === HttpEventType.UploadProgress) {
            if (event.total) {
              console.log('Upload Progress: ' + Math.round(event.loaded / event.total * 100) + '%');
            }
          } else if (event.type === HttpEventType.Response) {
            console.log(event);
          }
          console.log(event);
        })
    }
  }
}
