import {Component, ElementRef, ViewChild} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {Button, ButtonDirective} from "primeng/button";
import {Ripple} from "primeng/ripple";
import {Table, TableLazyLoadEvent, TableModule} from "primeng/table";
import {FormsModule} from "@angular/forms";
import {MultiSelectModule} from "primeng/multiselect";
import {DropdownModule} from "primeng/dropdown";
import {TagModule} from "primeng/tag";
import {SliderModule} from "primeng/slider";
import {ProgressBarModule} from "primeng/progressbar";
import {CurrencyPipe, DatePipe, NgClass} from "@angular/common";
import {InputTextModule} from "primeng/inputtext";
import {Customer} from "./models/customer";
import {Representative} from "./models/representative";
import {MysqlLazyTableComponent} from "./components/mysql-lazy-table/mysql-lazy-table.component";
import {PostgresLazyTableComponent} from "./components/postgres-lazy-table/postgres-lazy-table.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, Button, Ripple, ButtonDirective, TableModule, FormsModule, MultiSelectModule, DropdownModule, TagModule, SliderModule, ProgressBarModule, NgClass, DatePipe, CurrencyPipe, InputTextModule, MysqlLazyTableComponent, PostgresLazyTableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {

}
