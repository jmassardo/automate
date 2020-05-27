import { async, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { MockComponent } from 'ng2-mock-component';
import { StoreModule } from '@ngrx/store';
import {
  ngrxReducers,
  defaultInitialState,
  runtimeChecks,
  defaultRouterState,
  defaultRouterRouterState
} from 'app/ngrx.reducers';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';
import { DataBagsDetailsComponent } from './data-bags-details.component';


const declarations: any[] = [
  MockComponent({ selector: 'chef-heading' }),
  MockComponent({ selector: 'chef-icon' }),
  MockComponent({ selector: 'chef-loading-spinner' }),
  MockComponent({ selector: 'mat-select' }),
  MockComponent({ selector: 'mat-option' }),
  MockComponent({ selector: 'chef-page-header' }),
  MockComponent({ selector: 'chef-subheading' }),
  MockComponent({ selector: 'chef-toolbar' }),
  MockComponent({ selector: 'a', inputs: ['routerLink'] }),
  MockComponent({ selector: 'input', inputs: ['resetOrigin'] }),
  DataBagsDetailsComponent
];
const serverId = '6e98f609-586d-4816-a6de-e841e659b11d';
const orgId = '6e98f609-586d-4816-a6de';
const dataBagName = 'demo_data_bag';

describe('DataBagsDetailsComponent', () => {

  const initialState = {
    ...defaultInitialState,
    router: {
      ...defaultRouterState,
      state: {
        ...defaultRouterRouterState,
        url: `infrastructure/chef-servers/${serverId}/org/${orgId}`,
        params: { id: serverId, orgid: orgId, name: dataBagName }
      }
    }
  };

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: declarations,
      providers: [
        FeatureFlagsService
      ],
      imports: [
        RouterTestingModule,
        StoreModule.forRoot(ngrxReducers, { initialState, runtimeChecks })
      ]
    })
      .compileComponents();
  }));

  // it('should be created', () => {
  //   expect(component).toBeTruthy();
  // });

});
