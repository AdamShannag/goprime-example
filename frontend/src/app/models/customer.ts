import {Representative} from "./representative";
import {Country} from "./country";

export type Customer = {
  id?: number;
  name?: string;
  country?: Country;
  company?: string;
  date?: string;
  status?: string;
  activity?: number;
  representative?: Representative;
}

export type CustomersListResponse = {
  customers: Customer[];
  totalRecords: number;
}
