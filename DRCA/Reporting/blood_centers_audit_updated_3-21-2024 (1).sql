-- Quarterly Blood Centers KPIs Audit: Updated 3/21/2024


SELECT DISTINCT CASE WHEN site.site_name LIKE 'BA Downtown%'
                         THEN 'Blood Assurance Tennessee Downtown Donor Center'
                     WHEN site.site_name ILIKE 'BA Bellevue%'
                         THEN 'Blood Assurance Tennessee Bellevue Donor Center'
                     WHEN site.site_name ILIKE 'BA West End%'
                         THEN 'Blood Assurance Tennessee West End Donor Center'
                     WHEN site.site_name ILIKE 'BWNW%'
                         THEN BTRIM(REPLACE(site.site_name, 'BWNW', 'Bloodworks Northwest'))
                     WHEN site.site_name LIKE 'Cascade%'
                         THEN 'Cascade Regional Blood Services'
                     WHEN site.site_name ILIKE 'San Diego Blood Bank Mobile Diversion Pouch (SDBB Mobile DP) '
                         THEN 'San Diego Blood Bank Mobile'
                     WHEN site.site_name ILIKE 'STBTC Pavilion%'
                         THEN 'South Texas Blood & Tissue Center'
                     WHEN site.site_name ILIKE 'SC Lakewood%'
                         THEN 'SunCoast Blood Center'
                     WHEN site.google_group ILIKE '%scbmobile%'
                         THEN 'SunCoast Blood Center Mobile'
                    WHEN site.site_name ILIKE '%BA Clarksville%'
                        THEN 'Blood Assurance  Tennessee Clarksville'
                     WHEN site.google_group = 'hpo-site-bamobile'
                         THEN 'Blood Assurance Tennessee Mobile'
                    WHEN site.site_name ILIKE '%LifeSouth%'
                         THEN 'LifeSouth Community Blood Centers'
                    WHEN site.site_name ILIKE '%BA Cartersville%'
                        THEN 'BA Cartersville'
                     WHEN site.site_name ILIKE '%NYBC RIBC Providence%'
                          THEN 'Rhode Island Blood Center'
                     WHEN site.site_name LIKE '%(%'
                         THEN BTRIM(LEFT(INITCAP(site.site_name), STRPOS(site.site_name, '(') - 1))
                     ELSE 'Unknown'
                END                                                                                                     AS blood_center_name
                , COUNT(DISTINCT CASE WHEN bbo.bbo_biobank_order_id IS NOT NULL
                                      AND bbo.bbo_finalized_status = 'FINALIZED'
                                      AND bbo.bbo_collection_method = 'ON_SITE'
                                      AND bbs.bbs_baseline_test = 1
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_biosamples
                , COUNT(DISTINCT CASE WHEN bbo.bbo_finalized_status = 'FINALIZED'
                                      AND bbo.bbo_collection_method = 'ON_SITE'
                                      AND bbs.bbs_baseline_test = 1
                                      AND bbs.bbs_test = '1ED04'
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_4mLEDTA
                , COUNT(DISTINCT CASE WHEN bbo.bbo_finalized_status = 'FINALIZED'
                                      AND bbo.bbo_collection_method = 'ON_SITE'
                                      AND bbs.bbs_baseline_test = 1
                                      AND meas.code_value IN ('height', 'weight')
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_biosamplespm
                , COUNT(DISTINCT CASE WHEN bbs.bbs_test IN ('1ED04', '2ED04', '1ED10', '2ED10', '1HEP4',
                                                            '1PST8', '2PST8', '1SST8', '2SST8')
                                      AND bbo.bbo_tests_stored >= 4
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_atleast4blood
                , COUNT(DISTINCT CASE WHEN bbs.bbs_collected IS NOT NULL
                                      AND bbo.bbo_finalized_status = 'FINALIZED'
                                      AND bbo.bbo_collection_method = 'ON_SITE'
                                      AND bbs.bbs_baseline_test = 1
                                      AND bbs.bbs_test = '1UR10'
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_urine
                , COUNT(DISTINCT CASE WHEN bbo.bbo_finalized_status = 'FINALIZED'
                                      AND bbo.bbo_collection_method = 'ON_SITE'
                                      AND (CASE WHEN site.site_name LIKE '%Diversion Pouch%'
                                                    THEN 1
                                            ELSE 0 END) = 1
                                          THEN bbs.participant_id
                ELSE NULL END)                                                                                          AS participants_diversionpouch
FROM pdr.mv_participant_biobank_sample bbs
         LEFT JOIN pdr.mv_participant_biobank_order bbo ON (bbo.bbo_biobank_order_id = bbs.bbo_biobank_order_id)
         LEFT JOIN pdr.mv_site site ON (site.site_id = bbo.bbo_collected_site_id)
         LEFT JOIN pdr.mv_hpo hpo ON (hpo.hpo_id = site.hpo_id)
         LEFT JOIN pdr.mv_participant_pm pm ON (pm.participant_id = bbs.participant_id)
         LEFT JOIN pdr.mv_participant_measurement meas ON (meas.physical_measurements_id = pm.pm_physical_measurements_id)
         LEFT JOIN pdr.mv_participant_display pd ON (pd.participant_id = bbs.participant_id)
         LEFT JOIN pdr.mv_participant_retention_metrics ret ON (ret.participant_id = bbs.participant_id)
WHERE (meas.code_value IN ('height', 'weight')
       OR meas.code_value IS NULL)
    AND (site.google_group LIKE 'hpo-site-babellevue%'
        OR site.google_group = 'hpo-site-badowntown'
        OR site.google_group = 'hpo-site-badowntowndp_x3z8'
        OR site.google_group LIKE 'hpo-site-bawe%'
        OR site.google_group LIKE 'hpo-site-bwnw%'
        OR site.google_group LIKE 'hpo-site-cascadetacoma%'
        OR site.google_group LIKE 'hpo-site-sandiego%' OR site.google_group LIKE 'hpo-site-sdbb%'
        OR site.google_group LIKE 'hpo-site-stbtcpavilion%'
        OR site.google_group ILIKE 'hpo-site-scblakewood%'
        OR site.google_group ILIKE '%lifesouth%'
        OR site.google_group ILIKE '%baclarksville%'
        OR site.google_group ILIKE '%scbmobile%'
        OR site.google_group = 'hpo-site-bamobile'
        OR site.google_group ILIKE '%bacartersville%'
        OR site.google_group  = 'hpo-site-nybceribcprovidence')
    AND hpo.organization_type_id = 3
    AND bbo.test_participant = 0
    AND site.site_name IS NOT NULL
    AND bbs.bbs_collected IS NOT NULL
GROUP BY blood_center_name
ORDER BY blood_center_name ASC;