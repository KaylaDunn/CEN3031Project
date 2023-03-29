import { HttpClient, HttpEventType } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import * as _ from 'cypress/types/lodash';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  postData: any = {
    longitude: 0,
    latitude: 0
  }
  ngOnInit(): void {
    
  }

  selectedFile: File | null = null;

  constructor(private http: HttpClient) { }

  onFileSelected(event: any) {
    this.selectedFile = <File>event?.target.files[0];
  }

  onUpload() {
    this.http.post('0.0.0.0:3000/api/auth/createpost', this.postData, {
      reportProgress: true
    })

    if (this.selectedFile) {
      const fd = new FormData()
      this.http.post('0.0.0.0:3000/api/auth/createpost', fd, {
        reportProgress: true,
        observe: 'events'
      })

      fd.append('image', this.selectedFile, this.selectedFile.name);      
      
      this.http.post('0.0.0.0:3000/api/auth/addImageToPost/[id]', fd, { // url to backend function that accepts foreign data
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
