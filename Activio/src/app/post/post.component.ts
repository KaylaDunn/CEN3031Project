import { HttpClient, HttpEventType, HttpHeaders } from '@angular/common/http'; // allows for http requests
import { Component, Inject, OnInit } from '@angular/core';
import * as _ from 'cypress/types/lodash';

// defines PostComponent
@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})

export class PostComponent implements OnInit {

  // postData object, these values will be sent to backend API using POST request
  postData: any = {
    postDescription: "",
    longitude: 1.0,
    latitude: 1.0,
    locationname: ""
  }

  // lifecycle hook called when component is initialized; does nothing
  ngOnInit(): void {}

  // stores the user-uploaded file
  selectedFile: File | null = null;

  constructor(private http: HttpClient) { }

  // called when user selects a file
  onFileSelected(event: any) {  
    // sets selectedFile
    this.selectedFile = <File>event?.target.files[0];
  }

  // called when user hits Post
  onPost() {

    // if a file is selected, create FormData object
    if (this.selectedFile instanceof File) {
      
      // create request headers
      const headers = new HttpHeaders({
        'Content-Type': 'application/json'
      })

      // (1) POST request to the backend with postData object, creates new post
      // with headers
      this.http.post<any>('http://localhost:3000/api/auth/createpost', this.postData, {
        headers: headers,
        reportProgress: true,
        observe: 'response', // to get the post id
        withCredentials: true // authentication
      })
      
      .subscribe(response => {
        // get id from response body
        const id = response.body?.post.ID;
        console.log('id: ', id);

        // create FormData object
        const fd = new FormData();

        // appends the selected file to the FormData object
        fd.append('images', this.selectedFile!, this.selectedFile!.name);      

        // (2) another request to backend API with this object (upload/link image)
        // use backticks instead of quotes for id string formatting
        this.http.put(`http://localhost:3000/api/auth/addImageToPost/${id}`, fd, { // url to backend function that accepts foreign data
          // used to get progress events and response events from API
          reportProgress: true,
          observe: 'events',
          withCredentials: true // authentication
        }) 

        // called to subcribe the progress and response events from the API
        .subscribe(event => {
            if (event.type === HttpEventType.UploadProgress) {
              if (event.total) {
                console.log('Upload Progress: ' + Math.round(event.loaded / event.total * 100) + '%');
              }
            } else if (event.type === HttpEventType.Response) {
              console.log(event);
            }
        }, error => {
          console.error('Error uploading image: ', error);
        })
      }, error => {
        console.error('Error creating post: ', error);
      })
    }
  }
}
