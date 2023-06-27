# 🇲🇽 GOBL Mexico Tax Regime

Mexico uses the CFDI (Comprobante Fiscal Digital por Internet) format for their e-invoicing system.

## Public Documentation

* [Formato de factura (Anexo 20)](http://omawww.sat.gob.mx/tramitesyservicios/Paginas/anexo_20.htm)


## Local Codes

### `UsoCFDI`

The CFDI’s `UsoCFDI` field specifies how the invoice's recipient will use the invoice to deduce taxes for the expenditure made. The following table list all the supported values and how GOBL will map them from the invoice's tax tags:

| Code | Name | GOBL Tax Tag |
| --- | --- | --- |
| G01 | Adquisición de mercancías | `use+goods-acquisition` |
| G02 | Devoluciones, descuentos o bonificaciones | `use+returns` |
| G03 | Gastos en general | `use+general-expenses` |
| I01 | Construcciones | `use+construction` |
| I02 | Mobiliario y equipo de oficina por inversiones | `use+office-equipment` |
| I03 | Equipo de transporte | `use+transport-equipment` |
| I04 | Equipo de computo y accesorios | `use+computer-equipment` |
| I05 | Dados, troqueles, moldes, matrices y herramental | `use+manufacturing-tooling` |
| I06 | Comunicaciones telefónicas | `use+telephone-comms` |
| I07 | Comunicaciones satelitales | `use+satellite-comms` |
| I08 | Otra maquinaria y equipo | `use+other-machinery` |
| D01 | Honorarios médicos, dentales y gastos hospitalarios | `use+medical-expenses` |
| D02 | Gastos médicos por incapacidad o discapacidad | `use+medical-expenses+disability` |
| D03 | Gastos funerales | `use+funeral-expenses` |
| D04 | Donativos | `use+donation` |
| D05 | Intereses reales efectivamente pagados por créditos hipotecarios (casa habitación) | `use+mortgage-interest` |
| D06 | Aportaciones voluntarias al SAR | `use+sar-contribution` |
| D07 | Primas por seguros de gastos médicos | `use+medical-insurance` |
| D08 | Gastos de transportación escolar obligatoria | `use+school-transportation` |
| D09 | Depósitos en cuentas para el ahorro, primas que tengan como base planes de pensiones | `use+savings-deposit` |
| D10 | Pagos por servicios educativos (colegiaturas) | `use+school-fees` |
| S01 | Sin efectos fiscales | `use+no-tax-effects` |
| CP01 | Pagos | `use+suplementary-payment` |
| CN01 | Nómina | `use+payroll` |

### `FormaPago`

The CFDI’s `FormaPago` field specifies an invoice's means of payment. The following list list all the supported values and how GOBL will map them from the invoice's payment instructions key:

| Code | Name | GOBL Payment Instructions Key |
| --- | --- | --- |
| 01 | Efectivo | `cash` |
| 02 | Cheque nominativo | `cheque` |
| 03 | Transferencia electrónica de fondos | `credit-transfer` |
| 04 | Tarjeta de crédito | `card` |
| 05 | Monedero electrónico | `online+wallet` |
| 06 | Dinero electrónico | `online` |
| 08 | Vales de despensa | `other+grocery-vouchers  ` |
| 12 | Dación en pago | `other+in-kind` |
| 13 | Pago por subrogación | `other+subrogation` |
| 14 | Pago por consignación | `other+consignment` |
| 15 | Condonación | `other+debt-relief` |
| 17 | Compensación | `netting` |
| 23 | Novación | `other+novation` |
| 24 | Confusión | `other+merger` |
| 25 | Remisión de deuda | `other+remission` |
| 26 | Prescripción o caducidad | `other+expiration` |
| 27 | A satisfacción del acreedor | `other+extinguishment` |
| 28 | Tarjeta de débito | `card+debit` |
| 29 | Tarjeta de servicios | `card+services` |
| 30 | Aplicación de anticipos | `other+advance` |
| 31 | Intermediario pagos | `other+intermediary` |
| 99 | Por definir | `other` |

