import { TestBed } from '@angular/core/testing';
import { PostComponent } from './post.component'
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('PostComponent', () => {

  /*beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule], 
        providers: [AuthService]
      }));*/

  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule]
  }));
  
  it('should mount', () => {
    cy.mount(PostComponent)
  })

  it('buttons enabled except post', () => {
    cy.mount(PostComponent)
    cy.get('button[type="homeBtn"]').should('be.enabled')
    cy.get('button[type="btnUp"]').should('be.enabled')
  })

  it('fills inputs', () => {
    cy.mount(PostComponent)
    cy.get('input[name="loc"]').type('Gainesville')
    cy.get('input[name="bus"]').type('Krishna @ Plaza')
    cy.get('input[name="act"]').type('Getting Krishna')
  })

  // TODO:
  // it('post button disabled until input', () => {
  //   cy.mount(PostComponent)
  //   cy.get('button[name="postBtn"]').should('not.be.enabled')
  // })
})