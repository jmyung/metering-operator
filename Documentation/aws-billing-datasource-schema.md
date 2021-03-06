The schema below was obtained from Presto by running `describe datasource_aws_billing;` via Presto-cli.

```
                          Column                           |   Type    |     Extra     | Comment
-----------------------------------------------------------+-----------+---------------+---------
 identity_lineitemid                                       | varchar   |               |
 identity_timeinterval                                     | varchar   |               |
 bill_invoiceid                                            | varchar   |               |
 bill_billingentity                                        | varchar   |               |
 bill_billtype                                             | varchar   |               |
 bill_payeraccountid                                       | varchar   |               |
 bill_billingperiodstartdate                               | varchar   |               |
 bill_billingperiodenddate                                 | varchar   |               |
 lineitem_usageaccountid                                   | varchar   |               |
 lineitem_lineitemtype                                     | varchar   |               |
 lineitem_usagestartdate                                   | timestamp |               |
 lineitem_usageenddate                                     | timestamp |               |
 lineitem_productcode                                      | varchar   |               |
 lineitem_usagetype                                        | varchar   |               |
 lineitem_operation                                        | varchar   |               |
 lineitem_availabilityzone                                 | varchar   |               |
 lineitem_resourceid                                       | varchar   |               |
 lineitem_usageamount                                      | varchar   |               |
 lineitem_normalizationfactor                              | varchar   |               |
 lineitem_normalizedusageamount                            | varchar   |               |
 lineitem_currencycode                                     | varchar   |               |
 lineitem_unblendedrate                                    | varchar   |               |
 lineitem_unblendedcost                                    | varchar   |               |
 lineitem_blendedrate                                      | varchar   |               |
 lineitem_blendedcost                                      | double    |               |
 lineitem_lineitemdescription                              | varchar   |               |
 lineitem_taxtype                                          | varchar   |               |
 product_productname                                       | varchar   |               |
 product_accountassistance                                 | varchar   |               |
 product_alarmtype                                         | varchar   |               |
 product_architecturalreview                               | varchar   |               |
 product_architecturesupport                               | varchar   |               |
 product_availability                                      | varchar   |               |
 product_availabilityzone                                  | varchar   |               |
 product_bestpractices                                     | varchar   |               |
 product_cacheengine                                       | varchar   |               |
 product_caseseverityresponsetimes                         | varchar   |               |
 product_clockspeed                                        | varchar   |               |
 product_contenttype                                       | varchar   |               |
 product_currentgeneration                                 | varchar   |               |
 product_customerserviceandcommunities                     | varchar   |               |
 product_databaseengine                                    | varchar   |               |
 product_dedicatedebsthroughput                            | varchar   |               |
 product_deploymentoption                                  | varchar   |               |
 product_description                                       | varchar   |               |
 product_durability                                        | varchar   |               |
 product_ebsoptimized                                      | varchar   |               |
 product_ecu                                               | varchar   |               |
 product_endpointtype                                      | varchar   |               |
 product_enginecode                                        | varchar   |               |
 product_enhancednetworkingsupported                       | varchar   |               |
 product_executionfrequency                                | varchar   |               |
 product_executionlocation                                 | varchar   |               |
 product_feecode                                           | varchar   |               |
 product_feedescription                                    | varchar   |               |
 product_frequencymode                                     | varchar   |               |
 product_fromlocation                                      | varchar   |               |
 product_fromlocationtype                                  | varchar   |               |
 product_group                                             | varchar   |               |
 product_groupdescription                                  | varchar   |               |
 product_includedservices                                  | varchar   |               |
 product_instancefamily                                    | varchar   |               |
 product_instancetype                                      | varchar   |               |
 product_instancetypefamily                                | varchar   |               |
 product_io                                                | varchar   |               |
 product_launchsupport                                     | varchar   |               |
 product_licensemodel                                      | varchar   |               |
 product_location                                          | varchar   |               |
 product_locationtype                                      | varchar   |               |
 product_maxiopsburstperformance                           | varchar   |               |
 product_maxiopsvolume                                     | varchar   |               |
 product_maxthroughputvolume                               | varchar   |               |
 product_maxvolumesize                                     | varchar   |               |
 product_memory                                            | varchar   |               |
 product_memorygib                                         | varchar   |               |
 product_minvolumesize                                     | varchar   |               |
 product_networkperformance                                | varchar   |               |
 product_normalizationsizefactor                           | varchar   |               |
 product_operatingsystem                                   | varchar   |               |
 product_operation                                         | varchar   |               |
 product_operationssupport                                 | varchar   |               |
 product_origin                                            | varchar   |               |
 product_physicalprocessor                                 | varchar   |               |
 product_preinstalledsw                                    | varchar   |               |
 product_proactiveguidance                                 | varchar   |               |
 product_processorarchitecture                             | varchar   |               |
 product_processorfeatures                                 | varchar   |               |
 product_productfamily                                     | varchar   |               |
 product_programmaticcasemanagement                        | varchar   |               |
 product_provisioned                                       | varchar   |               |
 product_recipient                                         | varchar   |               |
 product_region                                            | varchar   |               |
 product_requestdescription                                | varchar   |               |
 product_requesttype                                       | varchar   |               |
 product_resourceendpoint                                  | varchar   |               |
 product_routingtarget                                     | varchar   |               |
 product_routingtype                                       | varchar   |               |
 product_servicecode                                       | varchar   |               |
 product_servicename                                       | varchar   |               |
 product_sku                                               | varchar   |               |
 product_storage                                           | varchar   |               |
 product_storageclass                                      | varchar   |               |
 product_storagemedia                                      | varchar   |               |
 product_technicalsupport                                  | varchar   |               |
 product_tenancy                                           | varchar   |               |
 product_thirdpartysoftwaresupport                         | varchar   |               |
 product_tolocation                                        | varchar   |               |
 product_tolocationtype                                    | varchar   |               |
 product_training                                          | varchar   |               |
 product_transfertype                                      | varchar   |               |
 product_usagefamily                                       | varchar   |               |
 product_usagetype                                         | varchar   |               |
 product_vcpu                                              | varchar   |               |
 product_version                                           | varchar   |               |
 product_volumetype                                        | varchar   |               |
 product_whocanopencases                                   | varchar   |               |
 pricing_publicondemandcost                                | varchar   |               |
 pricing_publicondemandrate                                | varchar   |               |
 pricing_term                                              | varchar   |               |
 pricing_unit                                              | varchar   |               |
 reservation_normalizedunitsperreservation                 | varchar   |               |
 reservation_totalreservednormalizedunits                  | varchar   |               |
 reservation_totalreservedunits                            | varchar   |               |
 reservation_unitsperreservation                           | varchar   |               |
 product_messagedeliveryfrequency                          | varchar   |               |
 product_messagedeliveryorder                              | varchar   |               |
 product_queuetype                                         | varchar   |               |
 reservation_amortizedupfrontcostforusage                  | varchar   |               |
 reservation_amortizedupfrontfeeforbillingperiod           | varchar   |               |
 reservation_effectivecost                                 | varchar   |               |
 reservation_endtime                                       | varchar   |               |
 reservation_modificationstatus                            | varchar   |               |
 reservation_recurringfeeforusage                          | varchar   |               |
 reservation_starttime                                     | varchar   |               |
 reservation_unusedamortizedupfrontfeeforbillingperiod     | varchar   |               |
 reservation_unusednormalizedunitquantity                  | varchar   |               |
 reservation_unusedquantity                                | varchar   |               |
 reservation_unusedrecurringfee                            | varchar   |               |
 reservation_upfrontvalue                                  | varchar   |               |
 lineitem_legalentity                                      | varchar   |               |
 product_capacitystatus                                    | varchar   |               |
 billing_period_start                                      | varchar   | partition key |
 billing_period_end                                        | varchar   | partition key |
```
