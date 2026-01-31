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
import { Options, LabelType } from "@angular-slider/ngx-slider";
import { AuthStorageService } from '../../services/auth-storage.service';
import { StoreService } from '../../services/store.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-store-page',
  standalone: false,
  templateUrl: './store-page.component.html',
  styles: ``
})
export class StorePageComponent implements AfterViewInit {

  totalProduct:number = 9;
  minValue: number = 0;
  maxValue: number = 100;
  errorMessage:string = "";
  categories:Array<any> = [];
  brands:Array<any> = [];
  products:Array<any> = [];
  bestSellers:Array<any> = [];
  topSellings:Array<any> = [];
  listProducts:Array<any> = [];
  loadingFilter:boolean = true;
  loadingProduct:boolean = true;
  selectedCategories: string[] = [];
  selectedBrands: string[] = [];
  orderBy:string = "published_at";
  sortBy:string = "desc";
  totalAll:number = 0;
  totalFiltered:number = 0;
  limit:number = 9;
  search:string = "";
  page:number = 1;
  authLogged:boolean = false;

  constructor(private authStorageService: AuthStorageService, private storeService: StoreService, private router: Router){}


  options: Options = {
    floor: 0,
    ceil: 500,
    translate: (value: number, label: LabelType): string => {
      switch (label) {
        case LabelType.Low:
          return "$" + value;
        case LabelType.High:
          return "$" + value;
        default:
          return "$" + value;
      }
    }
  };

  getCode(num:number): string{
    return (num).toString().padStart(2,"0")
  }

  ngAfterViewInit(): void {
    this.loadFilter();
    this.loadProduct();
    this.authLogged = this.authStorageService.getToken() !== null
  }

  loadFilter(): void{
    this.storeService.filter().subscribe({
        next: (res) => {
          setTimeout(() => {
            this.brands = res.brands
            this.categories = res.categories
            this.topSellings = res.tops
            this.maxValue = res.maxPrice
            this.minValue = res.minPrice
            this.loadingFilter = false
          }, 1500)
        },
        error: (err) => {
          const message = err.error.error || err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loadingFilter = false
        }
      });
  }

  loadProduct(): void{
    const fullQueryString = this.router.url.split('?')[1] ?? '';
    const result = fullQueryString ? '?' + fullQueryString : '';
    this.storeService.list(result).subscribe({
        next: (res) => {
          setTimeout(() => {
            this.totalAll = res.totalAll
            this.totalFiltered = res.totalFiltered
            this.limit = res.limit
            this.page = res.page
            this.listProducts = res.list
            this.loadingProduct = false
          }, 1500)
        },
        error: (err) => {
          const message = err.error.error || err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loadingProduct = false
        }
      });
  }

  handleFilter(event: any) {
    const e = event
    e.preventDefault();
    e.stopImmediatePropagation();

    let query = {}

    query = {
      ...query,
      minPrice: this.minValue,
      maxPrice: this.maxValue,
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

    if(this.selectedCategories.length > 0){
        query = {
          ...query,
          category: this.selectedCategories.join(",")
        }
    }

    if(this.selectedBrands.length > 0){
        query = {
          ...query,
          brand: this.selectedBrands.join(",")
        }
    }

    this.router.navigate([], {
      queryParams: query,
      //queryParamsHandling: 'merge',
    });
    this.loadingProduct = true

    setTimeout(() => {
      this.loadProduct()
    }, 1000)
  }

  handleCategory(item: string, event: any) {
    if (event.target.checked) {
      this.selectedCategories.push(item);
    } else {
      this.selectedCategories = this.selectedCategories.filter(i => i !== item);
    }
  }

  handleBrand(item: string, event: any) {
    if (event.target.checked) {
      this.selectedBrands.push(item);
    } else {
      this.selectedBrands = this.selectedBrands.filter(i => i !== item);
    }
  }

  handleLimit(value: number) {
    this.limit = value
  }

  handleOrderBy(value: string) {

    if(value === 'priceMax'){
      this.sortBy = "desc"
    }else if(value === 'priceMin'){
      this.sortBy = "asc"
    }

    this.orderBy = value
  }

  handleSearch(value: string) {
    this.search = value
  }

  loadPage(page: number) {
    this.page = page
    this.loadingProduct = true
    this.router.navigate([], {
      queryParams: { page: page },
      queryParamsHandling: 'merge',
    });
    setTimeout(() => {
      this.loadProduct()
    }, 1200)
  }

}
