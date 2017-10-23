import { Component, OnInit } from '@angular/core';
import { StatusService } from '../../services/status.service';

@Component({
  selector: 'app-status',
  templateUrl: './status.component.html',
  styleUrls: ['./status.component.css']
})
export class StatusComponent implements OnInit {

  status: any;

  constructor(private statusService: StatusService) {
    this.statusService.getAPIStatus()
    .subscribe(res => {
      this.status = res;
    })
  }
  ngOnInit() {
  }

}
