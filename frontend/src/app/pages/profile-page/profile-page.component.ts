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

import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../../services/profile.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { environment } from '../../../environments/environment.development';

@Component({
  selector: 'app-profile-page',
  standalone: false,
  templateUrl: './profile-page.component.html',
  styles: ``
})
export class ProfilePageComponent implements OnInit {

  loading: boolean = true
  upload: boolean = false
  submit: boolean = false
  activities:Array<any> = []
  formData: FormGroup;
  errorMessage:string = ""
  successMessage:string = ""
  userImage:string = "https://5an9y4lf0n50.github.io/demo-images/auth/user.png"

   genders = [
    { code: 'M', name: 'Male' },
    { code: 'F', name: 'Female' },
  ];

  constructor(private fb: FormBuilder, private profileService: ProfileService)
  {
     this.formData = this.fb.group({
      first_name: ['', [Validators.required]],
      last_name: ['', [Validators.required]],
      gender: ['', [Validators.required]],
      email: ['', [Validators.required]],
      phone: ['', [Validators.required]],
      city: ['', [Validators.required]],
      country: ['', [Validators.required]],
      zip_code: [''],
      address: ['', [Validators.required]]
    });
  }

  get first_name() {
    return this.formData.get('first_name');
  }

  get last_name() {
    return this.formData.get('last_name');
  }

  get gender() {
    return this.formData.get('gender');
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

  onFileChange(event: Event) {

    const input = event.target as HTMLInputElement;

    if (!input.files || input.files.length === 0) return;

    const file = input.files[0];
    const formData = new FormData();
    formData.append('file', file);

    this.upload = true
    this.errorMessage = ""
    this.successMessage = ""

    setTimeout(() => {
        this.profileService.upload(formData).subscribe({
        next: (res) => {
          const path = res.data
          this.upload = false
          this.errorMessage = ""
          this.successMessage = res.message
          this.loadActivity()
          this.userImage = (`${environment.apiUrl}/uploads/${path}`).toLowerCase()
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.upload = false
        }
      });
    }, 1500)

  }

  ngOnInit(): void {
      this.profileService.detail().subscribe({
        next: (res) => {
          const user = res

          if(user.image.String){
             this.userImage = this.profileService.getProfileImage(user.image.String)
          }

          delete user.id
          delete user.image
          delete user.status
          delete user.createdAt
          delete user.updatedAt
          this.formData.setValue(user)
        },
        error: (err) => {
          setTimeout(()=> {
            const message = err.error?.message || 'Something went wrong';
            this.errorMessage = message
          })
        }
      });
      this.loadActivity()
  }

  loadActivity(){
    this.loading = true
     this.profileService.activity(`limit=100`).subscribe({
        next: (res) => {
          setTimeout(async () => {
           this.activities = res
           this.loading = false
          }, 2000)
        },
        error: (err) => {
          setTimeout(()=> {
             const message = err.error?.message || 'Something went wrong';
            this.errorMessage = message
            this.loading = false
          })
        }
      });
  }

  onSubmit() {

    if (this.formData.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    this.submit = true
    this.errorMessage = ""
    this.successMessage = ""

    setTimeout(() => {
        this.profileService.update(this.formData.value).subscribe({
        next: (res) => {
          this.submit = false
          this.errorMessage = ""
          this.successMessage = res.message
          this.loadActivity()
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.submit = false
        }
      });
    }, 1500)

  }

}
