import {Component, ElementRef, ViewChild} from '@angular/core';
import {ButtonDirective} from "primeng/button";
import {CurrencyPipe, DatePipe, NgClass} from "@angular/common";
import {DropdownModule} from "primeng/dropdown";
import {InputTextModule} from "primeng/inputtext";
import {MultiSelectModule} from "primeng/multiselect";
import {PrimeTemplate} from "primeng/api";
import {ProgressBarModule} from "primeng/progressbar";
import {SliderModule} from "primeng/slider";
import {Table, TableLazyLoadEvent, TableModule} from "primeng/table";
import {CustomersListResponse} from "../../models/customer";
import {Representative} from "../../models/representative";
import {FormsModule} from "@angular/forms";
import {CustomerService} from "../../services/customer.service";
import {RepresentativeService} from "../../services/representative.service";
import {ListRequest} from "../../models/list.request";

@Component({
  selector: 'postgres-lazy-table',
  standalone: true,
  imports: [
    ButtonDirective,
    CurrencyPipe,
    DatePipe,
    DropdownModule,
    InputTextModule,
    MultiSelectModule,
    PrimeTemplate,
    ProgressBarModule,
    SliderModule,
    TableModule,
    FormsModule,
    NgClass
  ],
  templateUrl: './postgres-lazy-table.component.html',
  styleUrl: './postgres-lazy-table.component.scss'
})
export class PostgresLazyTableComponent {
  customersList: CustomersListResponse = {
    customers: [],
    totalRecords: 0,
  }
  representatives: Representative[] = [];

  statuses: any[] = [];
  activityValues: number[] = [0, 100];
  loading: boolean = true;

  @ViewChild('filter') filter!: ElementRef;

  constructor(private customerService: CustomerService, private representativeService: RepresentativeService) {
  }

  ngOnInit() {
    this.customerService.getPostgresCustomers().subscribe(res => {
      this.customersList.customers = res.customers;
      this.customersList.totalRecords = res.totalRecords;
      this.loading = false;
    });

    this.representativeService.getPostgresRepresentatives().subscribe(res => {
      this.representatives = res;
    })

    this.statuses = [
      {label: 'Unqualified', value: 'unqualified'},
      {label: 'Qualified', value: 'qualified'},
      {label: 'New', value: 'new'},
      {label: 'Negotiation', value: 'negotiation'},
      {label: 'Renewal', value: 'renewal'},
      {label: 'Proposal', value: 'proposal'}
    ];
  }

  onGlobalFilter(table: Table, event: Event) {
    table.filterGlobal((event.target as HTMLInputElement).value, 'contains');
  }

  clear(table: Table) {
    table.clear();
    this.filter.nativeElement.value = '';
  }

  lazyLoad($event: TableLazyLoadEvent) {
    this.loading = true
    this.customerService.getPostgresCustomers({
      filters: JSON.stringify($event.filters),
      first: $event.first,
      rows: $event.rows,
      sortField: $event.sortField ? $event.sortField : null,
      sortOrder: $event.sortOrder === 1 ? "ASC" : "DESC",
    } as ListRequest).subscribe(res => {
      this.customersList.customers = res.customers;
      this.customersList.totalRecords = res.totalRecords;
      this.loading = false;
    });

  }
}
