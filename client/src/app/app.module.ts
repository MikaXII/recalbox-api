import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { StatusComponent } from './component/status/status.component';
import {MaterialModule} from './modules/material/material.module'
import { StatusService } from './services/status.service';
import { HttpModule } from '@angular/http';

@NgModule({
  declarations: [
    AppComponent,
    StatusComponent
  ],
  imports: [
    BrowserModule,
    MaterialModule,
    HttpModule
  ],
  providers: [StatusService],
  bootstrap: [AppComponent]
})
export class AppModule { }
