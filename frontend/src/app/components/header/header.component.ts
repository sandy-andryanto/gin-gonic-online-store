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

import { AfterViewInit, Component, inject, OnDestroy, OnInit, TemplateRef } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { AuthStorageService } from '../../services/auth-storage.service';
import { ChangeDetectorRef } from '@angular/core';
import { ProfileService } from '../../services/profile.service';
import { HomeService } from '../../services/home.service';
import { SharedService } from '../../services/shared.service';
import { OrderService } from '../../services/order.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-header',
  standalone: false,
  templateUrl: './header.component.html',
  styles: ``
})
export class HeaderComponent implements AfterViewInit, OnInit, OnDestroy  {

  private modalService = inject(NgbModal);
  private readonly router = inject(Router);
  private subscription!: Subscription;

  constructor(
    private authStorageService: AuthStorageService,
    private cdr: ChangeDetectorRef,
    private profileService: ProfileService,
    private homeService: HomeService,
    private sharedService: SharedService,
    private orderService: OrderService
  ){}


  fliterSelected: number = 0
  auth: boolean = false
  txtSearch:string = ""
  userPhone:string = "Your Phone"
  userName:string = "Your Name"
  userLocation:string = "Your Location"
  errorMessage:string = ""
  categories:Array<any> = []
  carts:Array<any> = []
  order:any = {}
  whislists:Array<any> = []

  saveUserLogged(): void{
     this.profileService.detail().subscribe({
        next: (res) => {
          const usr = res
          this.authStorageService.saveUser(JSON.stringify(res))
          this.userPhone = usr.phone || "Your Phone"
          this.userName = `${usr.firstName} ${usr.lastName}` || "Your Phone"
          this.userLocation = `${usr.city}, ${usr.country}` || "Your Phone"
        },
        error: () => {}
      });
  }

  setFilter(event: any, index: number) {
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.fliterSelected = index
  }

  showModal(event:any, content: TemplateRef<any>) {
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.modalService.open(content)
  }

  redirectTo(event:any, route:string){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    setTimeout(() => {
      this.router.navigate([route]);
      this.modalService.dismissAll()
    })
  }

  ngAfterViewInit(): void {
    const token = this.authStorageService.getToken()
    const user = this.authStorageService.getUser()
    this.auth = token !== undefined && token !== null

    if(token !== undefined && token !== null){
      if(user === undefined || user === null){
        this.saveUserLogged()
      }else{
          const usr = JSON.parse(user)
          this.userPhone = usr.phone || "Your Phone"
          this.userName = `${usr.first_name} ${usr.last_name}` || "Your Phone"
          this.userLocation = `${usr.city}, ${usr.country}` || "Your Phone"
      }
    }

    this.homeService.component().subscribe({
      next: (res) => {
        this.categories = res.categories
      },
      error: (err) => {
        const message = err.error?.message || 'Something went wrong';
        this.errorMessage = message
      }
    });


    this.cdr.detectChanges()
    this.loadData();
  }

  logOut(event:any){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    this.authStorageService.removeUser()
    this.authStorageService.removeToken()
    setTimeout(() => {
       window.location.href = '/'
    }, 1500)
  }

  handleSearch(event:any){
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();
    if(this.categories[this.fliterSelected]){
      const category = this.categories[this.fliterSelected]
      this.router.navigate(['/store'], {
        queryParams: {
          category: category.id,
          search: this.txtSearch
        }
      });
    }else{
      this.router.navigate(['/store'], {
        queryParams: {
          search: this.txtSearch
        }
      });
    }
  }

   ngOnInit() {
    this.subscription = this.sharedService.loadData$.subscribe(() => {
      this.loadData();
    });
  }

  ngOnDestroy(): void {
   this.subscription.unsubscribe();
  }

   loadData() {

    const token = this.authStorageService.getToken()
    const user = this.authStorageService.getUser()

     if(token !== undefined && user !== null){
        this.orderService.session().subscribe({
          next: (res) => {
           this.carts = res.carts
           this.whislists = res.whislists
           this.order = res.order
          },
          error: (err) => {
            const message = err.error?.message || 'Something went wrong';
            this.errorMessage = message
          }
        });
     }


  }


}
