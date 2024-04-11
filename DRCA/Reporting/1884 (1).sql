SELECT distinct usr.user_id, family_name, given_name,workspace_id,rdr.name as Instutional_Affiliation, contact_email, rp_disease_of_focus,rp_intended_study,rp_anticipated_findings,rp_commercial_purpose,rp_ancestry,rp_disease_focused_research,rp_control_set,rp_drug_development,rp_educational,rp_ethics,rp_methods_development,rp_social_behavioral,rp_scientific_approach,usr.creation_time,access_tier_short_name,races
FROM rwb.mv_rwb_workspace wksp
    LEFT JOIN rwb.mv_rwb_users usr ON (usr.user_id = wksp.creator_id)
    LEFT JOIN rwb.mv_rdr_workspace rdr ON (rdr.workspace_source_id = wksp.workspace_id)
WHERE
     wksp.creation_time >= ('2020-05-27')
 and institution_id  not in (0,1)
    AND usr.institution_id IS NOT NULL
    AND usr.disabled = false
    AND rdr.exclude_from_public_directory = 0
    AND rdr.status = 'ACTIVE';
SELECT distinct usr.user_id, family_name, given_name,workspace_id,rdr.name as Instutional_Affiliation, contact_email, rp_disease_of_focus,rp_intended_study,rp_anticipated_findings,rp_commercial_purpose,rp_ancestry,rp_disease_focused_research,rp_control_set,rp_drug_development,rp_educational,rp_ethics,rp_methods_development,rp_social_behavioral,rp_scientific_approach,usr.creation_time,access_tier_short_name,races
FROM rwb.mv_rwb_workspace wksp
    LEFT JOIN rwb.mv_rwb_users usr ON (usr.user_id = wksp.creator_id)
    LEFT JOIN rwb.mv_rdr_workspace rdr ON (rdr.workspace_source_id = wksp.workspace_id)
WHERE
     wksp.creation_time <= ('2020-05-27')

    AND usr.institution_id IS NOT NULL
    AND usr.disabled = false
    AND rdr.exclude_from_public_directory = 0
    AND rdr.status = 'ACTIVE';
SELECT distinct usr.user_id, family_name, given_name,workspace_id,rdr.name as Instutional_Affiliation, contact_email, rp_disease_of_focus,rp_intended_study,rp_anticipated_findings,rp_commercial_purpose,rp_ancestry,rp_disease_focused_research,rp_control_set,rp_drug_development,rp_educational,rp_ethics,rp_methods_development,rp_social_behavioral,rp_scientific_approach,usr.creation_time,access_tier_short_name,races
FROM rwb.mv_rwb_workspace wksp
    LEFT JOIN rwb.mv_rwb_users usr ON (usr.user_id = wksp.creator_id)
    LEFT JOIN rwb.mv_rdr_workspace rdr ON (rdr.workspace_source_id = wksp.workspace_id)
WHERE
     wksp.creation_time >= ('2020-05-27')
 and institution_id  in (0,1)
    AND usr.institution_id IS NOT NULL
    AND usr.disabled = false
    AND rdr.exclude_from_public_directory = 0
    AND rdr.status = 'ACTIVE';