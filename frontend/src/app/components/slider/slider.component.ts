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

import { Component, Input, OnInit } from '@angular/core';
import { AuthStorageService } from '../../services/auth-storage.service';
import { OrderService } from '../../services/order.service';
import { SharedService } from '../../services/shared.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-slider',
  standalone: false,
  templateUrl: './slider.component.html',
  styles: ``
})
export class SliderComponent implements OnInit {

  constructor(
    private authStorageService: AuthStorageService,
    private orderService: OrderService,
    private sharedService: SharedService
  ){}

  @Input() products: any[] = []
  auth: boolean = true
  loading:boolean = false
  productId:number = 0
  errorMessage:string = ""

  slideConfig: any = {
    "slidesToShow": 3,
    "slidesToScroll": 1,
    "autoplay": true,
    "autoplaySpeed": 2000
  };

 ngOnInit(): void {
   this.auth = this.authStorageService.getToken() !== null
  }

  handleWhislist(event:any, product:number){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.productId = product
    this.loading = true
    this.orderService.wishlist(this.productId).subscribe({
        next: () => {
          setTimeout(() => {
            this.productId = 0
            this.loading = false
            Swal.fire({
              icon: 'success',
              title: 'Added to Wishlist',
              text: 'The item has been added to your wishlist.',
              timer: 1500,
              showConfirmButton: false,
            });
            this.sharedService.triggerLoadData()
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
