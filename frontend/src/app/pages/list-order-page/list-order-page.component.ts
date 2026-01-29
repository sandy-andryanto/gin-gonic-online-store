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
import { AuthStorageService } from '../../services/auth-storage.service';
import { OrderService } from '../../services/order.service';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { SharedService } from '../../services/shared.service';

@Component({
  selector: 'app-list-order-page',
  standalone: false,
  templateUrl: './list-order-page.component.html',
  styles: ``
})
export class ListOrderPageComponent implements AfterViewInit {

  loading:boolean = true
  orders:Array<any> = [];
  orderBy:string = "created_at";
  sortBy:string = "desc";
  totalAll:number = 0;
  totalFiltered:number = 0;
  limit:number = 10;
  search:string = "";
  page:number = 1;
  authLogged:boolean = false
  errorMessage:string = ""

  constructor(
    private authStorageService: AuthStorageService,
    private orderService: OrderService,
    private sharedService: SharedService,
    private router: Router
  ){}

   ngAfterViewInit(): void {
    this.loadData();
    this.authLogged = this.authStorageService.getToken() !== null
  }

   loadData(): void{
    const fullQueryString = this.router.url.split('?')[1] ?? '';
    const result = fullQueryString ? '?' + fullQueryString : '';
    this.orderService.list(result).subscribe({
        next: (res) => {
          setTimeout(() => {
            this.totalAll = res.totalAll
            this.totalFiltered = res.totalFiltered
            this.limit = res.limit
            this.page = res.page
            this.orders = res.list
            this.loading = false
          }, 1500)
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loading = false
        }
      });
  }

  handleLimit(event:any) {
    const value =  parseInt((event.target as HTMLSelectElement).value)
    this.limit = value
  }

  handleOrderBy(event:any) {
    const value =  (event.target as HTMLSelectElement).value
    this.orderBy = value
  }

  handleSortBy(event:any){
    const value =  (event.target as HTMLSelectElement).value
    this.sortBy = value
  }

  handleSearch(value:string) {
    this.search = value
  }

  handleFilter(event: any) {

    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();

    let query = {}

    query = {
      ...query,

      orderBy: this.orderBy,
      orderDir: this.sortBy,
      limit: this.limit,
      page: this.page
    }

    if(this.search){
       query = {
         ...query,
         search: this.search
       }
    }
    this.router.navigate([], {
      queryParams: query
    });
    this.loading = true

    setTimeout(() => {
      this.loadData()
    }, 1000)
  }

   loadPage(page: number) {
    this.page = page
    this.loading = true
    this.router.navigate([], {
      queryParams: { page: page },
      queryParamsHandling: 'merge',
    });
    setTimeout(() => {
      this.loadData()
    }, 1200)
  }

  async handleDelete(event:any, id:number){
     const e = event
      e.preventDefault();
      e.stopImmediatePropagation();

     const result = await Swal.fire({
        title: 'Are Your Sure ?',
        text: 'Please confirm that you accept these terms.?',
        icon: 'question',
        showCancelButton: true,
        confirmButtonText: 'Yes, Continue',
        cancelButtonText: 'Cancel',
        showLoaderOnConfirm: true,
        allowOutsideClick: () => !Swal.isLoading(),
        preConfirm: async () => {
          try {
            const res = await this.orderService.cancel(id).toPromise()
            return res;
          } catch (error:any) {
            if (error.error?.errors) {
                const messages = this.flattenErrors(error.error.errors);
                this.errorMessage = JSON.stringify(messages)
              }
          }
        }
      });

      if (result.isConfirmed) {
        Swal.fire({
          icon: 'success',
          title: 'Cancel Order Success',
          text: 'Your order been successfully canceled.'
        });
        this.sharedService.triggerLoadData()
        this.loadData()

    }

  }

  flattenErrors(errorObj: { [key: string]: string[] }): string[] {
    return Object.values(errorObj).flat();
  }

}
