import { CommentDialogComponent } from './comment-dialog.component'
import { HttpClientModule } from '@angular/common/http'
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog'

describe('CommentDialogComponent', () => {
  it('should mount', () => {
    cy.mount(CommentDialogComponent, {
      imports: [HttpClientModule, MatDialogModule],
      providers: [
        { provide: MAT_DIALOG_DATA, useValue: {} },
        { provide: MatDialogRef, useValue: {} }
      ]
    })
  })
})