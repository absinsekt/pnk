import { isUnset } from 'app/core/objects';
import { ListItem } from 'ui/lists/types';

export type DadataPayload = {
  area: string;
  area_fias_id: string;
  area_kladr_id: string;
  area_type: string;
  area_type_full: string;
  area_with_type: string;
  beltway_distance: string;
  beltway_hit: string;
  block: string;
  block_type: string;
  block_type_full: string;
  capital_marker: string;
  city: string;
  city_area: string;
  city_district: string;
  city_district_fias_id: string;
  city_district_kladr_id: string;
  city_district_type: string;
  city_district_type_full: string;
  city_district_with_type: string;
  city_fias_id: string;
  city_kladr_id: string;
  city_type: string;
  city_type_full: string;
  city_with_type: string;
  country: string;
  country_iso_code: string;
  federal_district: string;
  fias_actuality_state: string;
  fias_code: string;
  fias_id: string;
  fias_level: string;
  flat: string;
  flat_area: string;
  flat_price: string;
  flat_type: string;
  flat_type_full: string;
  geo_lat: string;
  geo_lon: string;
  geoname_id: string;
  history_values: string;
  house: string;
  house_fias_id: string;
  house_kladr_id: string;
  house_type: string;
  house_type_full: string;
  kladr_id: string;
  metro: string;
  okato: string;
  oktmo: string;
  postal_box: string;
  postal_code: string;
  qc: string;
  qc_complete: string;
  qc_geo: string;
  qc_house: string;
  region: string;
  region_fias_id: string;
  region_iso_code: string;
  region_with_type: string;
  settlement: string;
  settlement_fias_id: string;
  settlement_kladr_id: string;
  settlement_type: string;
  settlement_type_full: string;
  settlement_with_type: string;
  source: string;
  square_meter_price: string;
  street: string;
  street_fias_id: string;
  street_kladr_id: string;
  street_type: string;
  street_type_full: string;
  street_with_type: string;
  tax_office: string;
  tax_office_legal: string;
  timezone: string;
  unparsed_parts: string;
}

const DADATA_URL = 'https://suggestions.dadata.ru/suggestions/api/4_1/rs/';

export async function dadataSuggest(query: string): Promise<ListItem<DadataPayload>[]> {
  return fetch(`${DADATA_URL}suggest/address`, {
    method: 'POST',
    headers: {
      'Authorization': `Token ${PNK_DEV_DADATA_TOKEN}`,
      'Content-Type': 'application/json',
      'Accepte': 'application/json',
    },
    body: JSON.stringify({query, count: 10})
  })
  .then((data) => data.json())
  .catch(() => [])
  .then((res) => res.suggestions.map((s, idx) => ({id: idx, label: s.value, value: s.data})));
}

export function validHouse(errorMessage: string) {
  return function(value: DadataPayload): string {
    return isUnset(value) || isUnset(value.house) ? errorMessage : null;
  }
}

export function validApt(errorMessage: string) {
  return function(value: DadataPayload): string {
    return isUnset(value) || isUnset(value.flat) ? errorMessage : null;
  }
}
