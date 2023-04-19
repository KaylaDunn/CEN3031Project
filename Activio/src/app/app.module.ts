import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatMenuModule } from '@angular/material/menu';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatListModule } from '@angular/material/list';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTableModule } from '@angular/material/table';
import { MatBadgeModule } from '@angular/material/badge';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { HeaderComponent } from './components/header/header.component';
import { LoginComponent } from './login/login.component';
import { RouterModule } from '@angular/router';
import { WelcomeComponent } from './welcome/welcome.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { TextFieldModule } from '@angular/cdk/text-field';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { LogsuccessComponent } from './logsuccess/logsuccess.component';
import { CreateaccountComponent } from './createaccount/createaccount.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AuthService } from './auth.service';
import { PostComponent } from './post/post.component';
import { LogoutComponent } from './logout/logout.component';
import { CreateAccountVerifComponent } from './create-account-verif/create-account-verif.component';
import { UserProfileComponent } from './user-profile/user-profile.component';


import { LocationSearchComponent } from './location-search/location-search.component';
import { EditProfileInfoComponent } from './edit-profile-info/edit-profile-info.component';

import { MatDialogModule } from '@angular/material/dialog';
import { CommentDialogComponent } from './comment-dialog/comment-dialog.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LoginComponent,
    WelcomeComponent,
    LogsuccessComponent,
    CreateaccountComponent,
    PostComponent,
    LogoutComponent,
    CreateAccountVerifComponent,
    UserProfileComponent,
    LocationSearchComponent,
    EditProfileInfoComponent,
    CommentDialogComponent,
  ],
  
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    BrowserAnimationsModule,
    MatSidenavModule,
    MatGridListModule,
    MatMenuModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatExpansionModule,
    MatListModule,
    MatToolbarModule,
    MatTableModule,
    MatBadgeModule,
    MatSnackBarModule,
    HttpClientModule,
    RouterModule.forRoot([
      
      {
        path: 'login',
        component: LoginComponent
      },
      {
        path: '',
        component: HeaderComponent
      },
      {
        path: 'logsuccess',
        component: LogsuccessComponent
      },
      {
        path: 'createaccount',
        component: CreateaccountComponent
      },
      {
        path: 'post',
        component: PostComponent
      },
      {
        path: 'create-account-verif',
        component: CreateAccountVerifComponent
      },
      {
        path: 'userProfile',
        component: UserProfileComponent
      },
      {
        path: 'editAccount',
        component: EditProfileInfoComponent
      },
      {
        path: 'logout',
        component: LogoutComponent
        }
      ]),
      MatFormFieldModule,
      MatInputModule,
      TextFieldModule,
      MatAutocompleteModule,
      MatDialogModule
  ],
  providers: [AuthService],
  bootstrap: [AppComponent]
})
export class AppModule { }
