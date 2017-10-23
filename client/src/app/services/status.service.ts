import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Response} from '@angular/http';
import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';


@Injectable()
export class StatusService {

  constructor(private http: Http) { }

  getAPIStatus(): Observable<any> {
    return this.http.get(environment.baseUrl + 'status')
    .map((response: Response) => {
      console.log(response.json())
      return response.json()
    })
  }
}
