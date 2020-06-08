select *
from (
         select first_name,
                cast(count(*) filter (where gender = 'male' or gender = 'erkek') as float) /
                cast(count(*) as float) "Probability_Male",
                cast(count(*) filter (where gender = 'female' or gender = 'kadın') as float) /
                cast(count(*) as float) "Probability_FeMale",
                count(*)                "Total"
         from profiles
         where gender <> '' and first_name <> ''
         group by first_name
     ) t
where t."Probability_FeMale" > 0
  and t."Probability_FeMale" < 1
and "Total"> 10