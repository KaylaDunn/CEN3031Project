import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { HttpClient, HttpEventType, HttpHeaders } from '@angular/common/http';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
@Component({
  selector: 'app-edit-profile-info',
  templateUrl: './edit-profile-info.component.html',
  styleUrls: ['./edit-profile-info.component.css']
})
export class EditProfileInfoComponent implements OnInit {
  
  UserData: any = {}
  ImgData: any = {}
  me: any = {}
  constructor(private http: HttpClient, private dialog: MatDialog, private _auth: AuthService,
    private _router: Router) { }

 
  ngOnInit() : void {
    this.http.get<any>('http://localhost:3000/api/auth/me', {
      reportProgress: true,
      observe: 'response', // to get the post id
      withCredentials: true // authentication
    })
      .subscribe(response => {
        this.me = response.body;
      }, error => {
        console.error('Getting Account Info');
      })
  }
  // stores the user-uploaded file
  fileSelected = false;
  selectedFile: File | null = null;
  // called when user selects a file
  onFileSelected(event: any) {
    // sets selectedFiles
    this.selectedFile = <File>event?.target.files[0];
    this.fileSelected = true;
    // if a file is selected, create FormData object
    if (this.selectedFile instanceof File) {

      
      console.log(this.selectedFile!.name);
      // (1) POST request to the backend with postData object, creates new post
      // with headers
      const fd = new FormData();
      fd.append('image', this.selectedFile!, this.selectedFile!.name);
      
      this.http.post<any>('http://localhost:3000/api/auth/setprofileimage', fd, {
        observe: 'response', // to get the post id
        withCredentials: true // authentication
      })

        .subscribe(response => {
          console.log("Response Received");
          console.log(response);
          this.ImgData = response;
        }, error => {
          console.error('Error creating post: ', error);
          console.log(error);
        })

    }
  }



  User() {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    })
    console.log(this.UserData);
    this.http.put(`http://localhost:3000/api/auth/updateuser`, this.UserData, {
      withCredentials: true,
      headers: headers,
    })
      .subscribe(
        res => {
          console.log(this.UserData.email);
          console.log(this.UserData.username);
          console.log(this.UserData.firstName)
          console.log(this.UserData.lastName)
          console.log(this.UserData.birthday)
          console.log(this.UserData.phoneNumber)
          this._router.navigate(['/user-profile']);
        },
        err => console.log(err)
      )
  }

}
