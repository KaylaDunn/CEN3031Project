import { HttpClientModule } from '@angular/common/http'
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog'
import { AddImageComponent } from './add-image.component'

describe('AddImageComponent', () => {
  it('should mount', () => {
    cy.mount(AddImageComponent, {
      imports: [HttpClientModule, MatDialogModule],
      providers: [
        { provide: MAT_DIALOG_DATA, useValue: {} },
        { provide: MatDialogRef, useValue: {} }
      ]
    })
  });
})