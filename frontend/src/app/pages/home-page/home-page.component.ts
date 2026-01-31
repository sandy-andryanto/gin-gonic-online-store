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

import { AfterViewInit, Component } from '@angular/core';
import { HomeService } from '../../services/home.service';
@Component({
  selector: 'app-home-page',
  standalone: false,
  templateUrl: './home-page.component.html',
  styles: ``
})
export class HomePageComponent implements AfterViewInit {

  loading: boolean = true
  categories:Array<any> = []
  products:Array<any> = []
  topSellings:Array<any> = []
  bestSellers:Array<any> = []
  activeTab1:number = 0
  activeTab2:number = 0
  errorMessage:string = ""

  constructor(private homeService: HomeService){}

  setTab1(event:any, index:number){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.activeTab1 = index;
  }

  setTab2(event:any, index:number){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.activeTab2 = index;
  }

  ngAfterViewInit(): void {
     this.homeService.page().subscribe({
        next: (res) => {
          setTimeout(() => {
            this.categories = res.categories
            this.products = res.products
            this.topSellings = res.topSellings
            this.bestSellers = res.bestSellers
            this.loading = false
          }, 1500)
        },
        error: (err) => {
          setTimeout(()=> {
            const message = err.error?.message || 'Something went wrong';
            this.errorMessage = message
          })
        }
      });
  }


}
