select ps.participant_id,first_name,middle_name,last_name,email,phone_number,ps.participant_origin from participant_summary ps
left join participant p  on p.participant_id  = ps.participant_id
where is_test_participant = 0  and is_ghost_id is null and p.hpo_id not in (21) and p.sign_up_time between date_sub(now(),INTERVAL 1 WEEK) and now();

