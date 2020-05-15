import { Component, Input, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { FormGroup, Validators, FormBuilder } from '@angular/forms';
import { Subject, combineLatest } from 'rxjs';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { LayoutFacadeService, Sidebar } from 'app/entities/layout/layout.facade';
import { filter, takeUntil } from 'rxjs/operators';
import { EntityStatus, pending } from 'app/entities/entities';
import { AdminKey } from 'app/entities/reset-admin-key/reset-admin-key.model';
import {
   updateStatus
} from 'app/entities/reset-admin-key/reset-admin-key.selectors';
import { UpdateAdminKey } from 'app/entities/reset-admin-key/reset-admin-key.actions';

@Component({
  selector: 'app-reset-admin-key',
  templateUrl: './reset-admin-key.component.html',
  styleUrls: ['./reset-admin-key.component.scss']
})

export class ResetAdminKeyComponent implements OnInit {
  @Input() serverId: string;
  @Input() orgId: string;

  public org: AdminKey;
  public saveSuccessful = false;
  public saveInProgress = false;
  public isLoading = true;
  public resetKeyForm: FormGroup;
  private isDestroyed = new Subject<boolean>();

  constructor(
    private fb: FormBuilder,
    private store: Store<NgrxStateAtom>,
    private layoutFacade: LayoutFacadeService
  ) {
    this.resetKeyForm = this.fb.group({
      admin_key: ['', [Validators.required]]
    });
  }

  ngOnInit() {
    this.layoutFacade.showSidebar(Sidebar.Infrastructure);

    combineLatest([
      this.store.select(updateStatus)
    ]).pipe(
      takeUntil(this.isDestroyed)
    ).subscribe(([updateSt]) => {
      this.isLoading = updateSt === EntityStatus.loading;
    });

    this.store.select(updateStatus).pipe(
      takeUntil(this.isDestroyed),
      filter(state => this.saveInProgress && !pending(state)))
      .subscribe((state) => {
        this.saveInProgress = false;
        this.saveSuccessful = (state === EntityStatus.loadingSuccess);
        if (this.saveSuccessful) {
          this.resetKeyForm.markAsPristine();
          this.resetKeyForm.reset();
        }
      });
  }

  saveNewKey(): void {
    this.saveSuccessful = false;
    this.saveInProgress = true;
    const admin_key: string = this.resetKeyForm.controls.admin_key.value.trim();
    this.store.dispatch(new UpdateAdminKey({
      server_id: this.serverId, org_id: this.orgId, admin_Key: {admin_key}
    }));
  }
}
