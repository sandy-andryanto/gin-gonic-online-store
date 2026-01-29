import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { ErrorPageComponent } from './pages/error-page/error-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { ConfirmPageComponent } from './pages/confirm-page/confirm-page.component';
import { ForgotPasswordPageComponent } from './pages/forgot-password-page/forgot-password-page.component';
import { ResetPasswordPageComponent } from './pages/reset-password-page/reset-password-page.component';
import { PasswordPageComponent } from './pages/password-page/password-page.component';
import { ProfilePageComponent } from './pages/profile-page/profile-page.component';
import { CartPageComponent } from './pages/cart-page/cart-page.component';
import { CheckoutPageComponent } from './pages/checkout-page/checkout-page.component';
import { StorePageComponent } from './pages/store-page/store-page.component';
import { ListOrderPageComponent } from './pages/list-order-page/list-order-page.component';
import { DetailOrderPageComponent } from './pages/detail-order-page/detail-order-page.component';

const routes: Routes = [
  { path: 'home', component: HomePageComponent },
  { path: 'error', component: ErrorPageComponent },
  { path: 'auth/login', component: LoginPageComponent },
  { path: 'auth/register', component: RegisterPageComponent },
  { path: 'auth/register/confirm/:token', component: ConfirmPageComponent },
  { path: 'auth/email/forgot', component: ForgotPasswordPageComponent },
  { path: 'auth/email/reset/:token', component: ResetPasswordPageComponent },
  { path: 'account/profile', component: ProfilePageComponent },
  { path: 'account/password', component: PasswordPageComponent },
  { path: 'order/list', component: ListOrderPageComponent },
  { path: 'order/detail/:id', component: DetailOrderPageComponent },
  { path: 'cart/:id', component: CartPageComponent },
  { path: 'checkout', component: CheckoutPageComponent },
  { path: 'store', component: StorePageComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: '**', redirectTo: '/error', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
