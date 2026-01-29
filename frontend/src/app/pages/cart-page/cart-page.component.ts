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

import { AfterContentInit, Component, inject } from '@angular/core';
import { OrderService } from '../../services/order.service';
import { SharedService } from '../../services/shared.service';
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
@Component({
  selector: 'app-cart-page',
  standalone: false,
  templateUrl: './cart-page.component.html',
  styles: ``
})
export class CartPageComponent implements AfterContentInit {

  maxStock:number = 0
  rating:number = 0
  price:number = 0
  tabActive:number = 2
  qty:number = 1
  total:number = 0
  errorMessage:string = ""
  authLogged:boolean = false
  loadingCart:boolean = true
  loadingReview:boolean = true
  images:Array<any> = [];
  productRelated:Array<any> = [];
  product:any = {}
  reviews:Array<any> = [];
  categories:Array<string> = [];
  sizes:Array<any> = [];
  inventories:Array<any> = [];
  colours:Array<any> = [];
  formData: FormGroup;
  sizeId:number = 0
  colourId:number = 0
  private readonly router = inject(Router);

  slideConfig: any = {
    "slidesToShow": 1,
    "slidesToScroll": 1,
    "autoplay": true,
    "autoplaySpeed": 2000
  };

  constructor(
    private fb: FormBuilder,
    private orderService: OrderService,
    private route: ActivatedRoute,
    private sharedService: SharedService
  ){
    this.formData = this.fb.group({
      name: ['', [Validators.required]],
      email: ['', [Validators.required]],
      review: ['', [Validators.required]]
    });
  }

   get name() {
    return this.formData.get('name');
  }

   get email() {
    return this.formData.get('email');
  }

   get review() {
    return this.formData.get('review');
  }

  handleTab(event:any, index:number) {
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.tabActive = index
  }


  ngAfterContentInit(): void {
    this.total = this.price
    this.loadProduct()
  }

  loadProduct() : void{
    const id = this.route.snapshot.paramMap.get('id') || '';
    this.loadingCart = true
    this.orderService.cart(parseInt(id)).subscribe({
        next: (res) => {
          setTimeout(() => {
              this.total = res.product.Price
              this.sizes = res.sizes
              this.colours = res.colours
              this.inventories = res.inventories
              this.images = res.images
              this.product = res.product
              this.productRelated = res.productRelated
              this.loadingCart = false
              this.price = res.product.Price
              this.categories = (res.product.CategoryName).split(",")
              this.formData.setValue({
                name: `${res.user.first_name.String} ${res.user.last_name.String}`,
                email: res.user.email,
                review: 'Your Review Here'
              });
              this.loadReview()
          }, 1500)
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loadingCart = false
        }
      });
  }

  loadReview() : void{
    const id = this.route.snapshot.paramMap.get('id') || '';
    this.loadingReview = true
    this.orderService.listReview(parseInt(id)).subscribe({
        next: (res) => {
          setTimeout(() => {
              this.reviews = res
              this.loadingReview = false
          }, 1500)
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loadingReview = false
        }
      });
  }

  setQty(event:any, type:string){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    if(this.sizeId > 0 && this.colourId > 0)
    {
      if (type === '+') {
         if(this.maxStock > 0 && this.maxStock === this.qty){
           this.qty = this.maxStock
           this.total = this.price * this.maxStock
         }else{
            this.qty = this.qty + 1
            this.total = this.qty * this.price
         }
      }else{
        if(this.qty > 1){
          this.qty = this.qty - 1
          this.total = this.qty * this.price
        }
      }
    }

  }

  handleWhislist(event:any){
      const id = this.route.snapshot.paramMap.get('id') || '';
      const e = event
      e.preventDefault();
      e.stopImmediatePropagation();
       Swal.fire({
          title: 'Please wait...',
          text: 'Adding to Wishlist...',
          allowOutsideClick: false,
          didOpen: () => {
            Swal.showLoading();
          }
        });
      this.orderService.wishlist(parseInt(id)).subscribe({
          next: () => {
            setTimeout(() => {
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

    handleRating(event:any){
      const e = event
      this.rating = e.rating
    }

    async onSubmitReview() {

      if (this.formData.invalid) {
        this.errorMessage = 'Form is invalid'
        return;
      }

      const id = this.route.snapshot.paramMap.get('id') || '';
      let formSubmit = this.formData.value
      formSubmit = {
         ...formSubmit,
         rating: this.rating
      }

       const result = await Swal.fire({
        title: 'Adding New Review',
        text: 'Are you sure you want to create new review?',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonText: 'Yes, Continue',
        cancelButtonText: 'Cancel',
        showLoaderOnConfirm: true,
        allowOutsideClick: () => !Swal.isLoading(),
        preConfirm: async () => {
          try {
            const res = await this.orderService.createReview(parseInt(id), formSubmit).toPromise();
            return res;
          } catch (error: any) {
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
          title: 'Review has been added!',
          text: 'Your review has been successfully added to this product.'
        });
        this.loadReview()
        this.rating = 0
      }

    }

    flattenErrors(errorObj: { [key: string]: string[] }): string[] {
      return Object.values(errorObj).flat();
    }

    handleInventory(){
      if(this.sizeId > 0 && this.colourId > 0){
        const getStock = this.inventories.filter(row => row.sizeId === this.sizeId && row.colourId === this.colourId)
        if(getStock.length > 0){
           const stock = getStock[0].stock
           this.maxStock = stock
        }
      }
    }

    handleSize(event:any){
      const e = event
      e.preventDefault();
      e.stopImmediatePropagation();
      this.sizeId = parseInt((event.target as HTMLSelectElement).value)
      this.handleInventory()
    }

    handleColour(event:any){
      const e = event
      e.preventDefault();
      e.stopImmediatePropagation();
      this.colourId = parseInt((event.target as HTMLSelectElement).value)
      this.handleInventory()
    }

    async handleAddCart(event:any){

      const e = event
      e.preventDefault();
      e.stopImmediatePropagation();
      const id = this.route.snapshot.paramMap.get('id') || '';

      const formSubmit = {
        size_id: this.sizeId,
        colour_id: this.colourId,
        price: this.price,
        qty: this.qty,
        total: this.total
      }

      const result = await Swal.fire({
          title: 'Adding Cart Confirmation',
          text: 'Are you sure you want to place your order?',
          icon: 'warning',
          showCancelButton: true,
          confirmButtonText: 'Yes, Continue',
          cancelButtonText: 'Cancel',
          showLoaderOnConfirm: true,
          allowOutsideClick: () => !Swal.isLoading(),
          preConfirm: async () => {
            try {
              const res = await this.orderService.createCart(parseInt(id), formSubmit).toPromise();
              return res;
            } catch (error:any) {
              const messages = this.flattenErrors(error.error.errors);
              this.errorMessage = JSON.stringify(messages)
            }
          }
        });

        if (result.isConfirmed) {
          Swal.fire({
            icon: 'success',
            title: 'Cart has been added!',
            text: 'Your product has been successfully added to cart.'
          });
          this.sharedService.triggerLoadData()
          setTimeout(() => {
            this.router.navigate([`order/list`]);
          }, 1500)
        }


    }

}
