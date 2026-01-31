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

import { AfterContentInit, Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OrderService } from '../../services/order.service';
import { SharedService } from '../../services/shared.service';

@Component({
  selector: 'app-checkout-page',
  standalone: false,
  templateUrl: './checkout-page.component.html',
  styles: ``
})
export class CheckoutPageComponent implements AfterContentInit {

  payment:number = 0
  accept:boolean = false
  loading:boolean = true
  formData: FormGroup;
  errorMessage:string = ""
  order:any = {}
  carts:Array<any> = [];
  payments:Array<any> = [];
  discount:number = 0
  shipment:number = 0
  taxes:number = 0
  user:any = {}
  private readonly router = inject(Router);

  constructor(
    private fb: FormBuilder,
    private orderService: OrderService,
    private sharedService: SharedService
  )
  {
     this.formData = this.fb.group({
      first_name: ['', [Validators.required]],
      last_name: ['', [Validators.required]],
      email: ['', [Validators.required]],
      phone: ['', [Validators.required]],
      city: ['', [Validators.required]],
      country: ['', [Validators.required]],
      zip_code: [''],
      address: ['', [Validators.required]],
      notes: ['']
    });
  }


  get first_name() {
    return this.formData.get('first_name');
  }

  get last_name() {
    return this.formData.get('last_name');
  }

  get email() {
    return this.formData.get('email');
  }

  get phone() {
    return this.formData.get('phone');
  }

  get city() {
    return this.formData.get('city');
  }

  get country() {
    return this.formData.get('country');
  }

  get zip_code() {
    return this.formData.get('zip_code');
  }

  get address() {
    return this.formData.get('address');
  }

  setAccept(event:any){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.accept = e.target.checked
  }

  setPayment(event:any, index:number){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.payment = index
  }

  ngAfterContentInit(): void {
    this.loading = true
    this.orderService.checkoutInitial().subscribe({
        next: (res) => {
          setTimeout(() => {
            let resOrder = res.order
            resOrder = {
              ...resOrder,
              total_taxes: parseFloat(resOrder.total_taxes).toFixed(2),
              total_paid: parseFloat(resOrder.total_paid).toFixed(2),
              total_discount: parseFloat(resOrder.total_discount).toFixed(2)
            }

            this.order = resOrder
            this.carts = res.carts
            this.payments = res.payments
            this.discount = parseFloat(res.discount)
            this.shipment = parseFloat(res.shipment)
            this.taxes = parseFloat(res.taxes)
            this.payment = res.order.payment_id
            this.formData.setValue({
              first_name: res.user.first_name.String,
              last_name:res.user.last_name.String,
              email: res.user.email,
              phone: res.user.phone,
              city: res.user.city.String,
              country: res.user.country.String,
              zip_code: res.user.zip_code.String,
              address: res.user.address.String,
              notes: "",
            });
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

  async onSubmit() {

      if (this.formData.invalid) {
        this.errorMessage = 'Form is invalid'
        return;
      }

      const result = await Swal.fire({
        title: 'Confirm Checkout ?',
        text: 'Are you sure you want to place your order?',
        icon: 'question',
        showCancelButton: true,
        confirmButtonText: 'Yes, Checkout',
        cancelButtonText: 'Cancel',
        showLoaderOnConfirm: true,
        allowOutsideClick: () => !Swal.isLoading(),
        preConfirm: async () => {
          try {

            let formData = this.formData.value
            formData = {
              ...formData,
              payment_id: this.payment
            }

            const res = await this.orderService.checkoutSubmit(formData).toPromise();
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
          title: 'Checkout Success',
          text: 'Your order been successfully checkout.'
        });
        this.sharedService.triggerLoadData()
        setTimeout(() => {
            this.router.navigate([`order/list`]);
        }, 1500)

    }
  }

   flattenErrors(errorObj: { [key: string]: string[] }): string[] {
      return Object.values(errorObj).flat();
    }



}
