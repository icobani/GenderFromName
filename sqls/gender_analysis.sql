select *,
       CASE
           WHEN t."probability_female" > t."probability_male" THEN 'female'
           else 'male'
           END "gender"
from (
         select first_name,
                cast(count(*) filter (where gender = 'male' or gender = 'erkek') as float) /
                cast(count(*) as float) "probability_male",
                cast(count(*) filter (where gender = 'female' or gender = 'kadın') as float) /
                cast(count(*) as float) "probability_female",
                count(*)                "total_count"
         from profiles
         where gender <> ''
           and first_name <> ''
         group by first_name
     ) t