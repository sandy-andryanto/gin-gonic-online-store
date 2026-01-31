/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.official@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class StoreService {

  constructor(private http: HttpClient) {}

   filter(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/shop/filter`);
   }

   list(param:string): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/shop/list${param}`);
   }

}
