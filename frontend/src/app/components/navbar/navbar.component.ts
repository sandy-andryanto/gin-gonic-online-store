/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.blade@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

import { AfterViewInit, Component } from '@angular/core';
import { HomeService } from '../../services/home.service';

@Component({
  selector: 'app-navbar',
  standalone: false,
  templateUrl: './navbar.component.html',
  styles: ``
})
export class NavbarComponent implements AfterViewInit {

   categories:Array<any> = []
   errorMessage:string = ""

  constructor(private homeService: HomeService){}

  ngAfterViewInit(): void {
    this.homeService.component().subscribe({
      next: (res) => {
        this.categories = res.categories
      },
      error: (err) => {
        const message = err.error?.message || 'Something went wrong';
        this.errorMessage = message
      }
    });
  }

}
