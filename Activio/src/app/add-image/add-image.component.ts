import { Component, Inject } from '@angular/core';
import { HttpClient, HttpEventType } from '@angular/common/http';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-add-image',
  templateUrl: './add-image.component.html',
  styleUrls: ['./add-image.component.css']
})
export class AddImageComponent {

  post: any;
  constructor(private http: HttpClient, @Inject(MAT_DIALOG_DATA) public data: { id: number }, private dialog: MatDialog, private dialogRef: MatDialogRef<AddImageComponent>) { }
  fileSelected = false;
  selectedFile: File | null = null;

  ngOnInit(): void {
    if (this.data) {
      // get post with that id
      this.http.get<any>(`http://localhost:3000/api/post/${this.data.id}`)
      .subscribe(response => {
  
        this.post = response.post;
        
      }, error => {
        console.error(error);
      });
    }
  }

  onFileSelected(event: any) {
    // sets selectedFile
    this.selectedFile = <File>event?.target.files[0];
    this.fileSelected = true;
  }

  onPost() {
    // if a file is selected, create FormData object
    if (this.selectedFile instanceof File) {
             
      // create FormData object
      const fd = new FormData();

      // appends the selected file to the FormData object
      fd.append('images', this.selectedFile!, this.selectedFile!.name);

      // request to backend API with this object (upload/link image)
      this.http.put(`http://localhost:3000/api/auth/addImageToPost/${this.post.id}`, fd, { // url to backend function that accepts foreign data
        // used to get progress events and response events from API
        reportProgress: true,
        observe: 'events',
        withCredentials: true // auth
      }) 

      // called to subcribe the progress and response events from the API
      .subscribe(event => {
          if (event.type === HttpEventType.UploadProgress) {
            if (event.total) {
              console.log('Upload Progress: ' + Math.round(event.loaded / event.total * 100) + '%');
            }
          } else if (event.type === HttpEventType.Response) {
            console.log(event);
            this.dialogRef.close();
        }
      }, error => {
        console.error('Error uploading image: ', error);
      })
    }
  }
}
