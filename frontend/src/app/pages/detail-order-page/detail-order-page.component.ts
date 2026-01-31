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
import { AuthStorageService } from '../../services/auth-storage.service';
import { OrderService } from '../../services/order.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-detail-order-page',
  standalone: false,
  templateUrl: './detail-order-page.component.html',
  styles: ``
})
export class DetailOrderPageComponent implements AfterViewInit {

  payment:any = {}
  carts:Array<any> = [];
  billings:Array<any> = [];
  order:any = {}
  loading:boolean = true
  errorMessage:string = ""
  discount:number = 0
  taxes:number = 0
  shipment:number = 0

  constructor(private authStorageService: AuthStorageService, private orderService: OrderService, private router: Router,  private route: ActivatedRoute,){}

  ngAfterViewInit(): void {
    const id = this.route.snapshot.paramMap.get('id') || '';
     this.orderService.detail(parseInt(id)).subscribe({
        next: (res) => {
           setTimeout(() => {
            this.carts = res.carts
            this.order = res.order
            this.billings = res.billings
            this.discount = res.discount
            this.taxes = res.taxes
            this.shipment = res.shipment
            this.payment = res.payment
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



}
