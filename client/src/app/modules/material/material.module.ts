import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatToolbarModule, MatCardModule, MatListModule} from '@angular/material';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

@NgModule({
  imports: [MatToolbarModule, BrowserAnimationsModule, MatCardModule,MatListModule],
  exports: [MatToolbarModule ,BrowserAnimationsModule, MatCardModule,MatListModule],
  declarations: []
})
export class MaterialModule { }
