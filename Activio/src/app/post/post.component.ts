import { HttpClient, HttpEventType, HttpHeaders } from '@angular/common/http'; // allows for http requests
import { Component, Inject, OnInit } from '@angular/core';
import * as _ from 'cypress/types/lodash';

// interpret id from json object
interface PostResponse {
  id: number;
}

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
    longitude: 0,
    latitude: 0,
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

    // debugging
    console.log('onPost called');

    // if a file is selected, create FormData object
    if (this.selectedFile instanceof File) {
      console.log('file selected');
      
      const fd = new FormData()

      // debugging: verify postData
      console.log('name: ', this.postData.locationname);
      console.log("act: ", this.postData.postDescription);

      // get cookie value
      // const ugid = document.cookie.match(/ugid=([^;]+)/)?.[1];
      console.log('cookies: ', document.cookie);

      // create request headers
      const headers = new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${ugid}'
      })


      // (1) POST request to the backend with postData object, creates new post
      // with headers
      this.http.post<PostResponse>('http://localhost:3000/api/auth/createpost', this.postData, {
        headers: headers,
        reportProgress: true,
        observe: 'response' // to get the post id
      })
      
      .subscribe(response => {
        // get id from response body
        const id = response.body?.id;
        
        // appends the selected file to the FormData object
        fd.append('image', this.selectedFile!, this.selectedFile!.name);      
        
        // (2) another POST request to backend API with this object (upload/link image)
        this.http.post('0.0.0.0:3000/api/auth/addImageToPost/${id}', fd, { // url to backend function that accepts foreign data
          // used to get progress events and response events from API
          reportProgress: true,
          observe: 'events'
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
            console.log(event);
        }, error => {
          console.error('Error uploading image: ', error);
        })
      }, error => {
        console.error('Error creating post: ', error);
      })
    }
  }
}
