package main

func getAphorismList() []string {
	s := []string{

		"Некоторое утверждают, что неравенство — это естественный результат рыночного обмена. Это очевидное заблуждение. Неравенство — это условие, при котором обмен вообще имеет смысл.",

		"Среди этих важных и серьезных забот казалось совершенно неуместным, что он хлопотал без всякого основательного повода, но исключительно ради популярности, о дешевизне продуктов; а дело это такое, что если неудачно его повести, то обычно в результате бывает голод. (C) Аммиан Марцеллин",

		"На протяжении всей истории, естественным состоянием человечества является бедность. Достижения, которые позволяют превысить эту норму, - там и сям, время от времени, - заслуга очень ограниченного меньшинства, которое часто презирают, обычно осуждают, и которому практически всегда противодействуют все здравомыслящие люди. Когда этому меньшинству не дают созидать или (как иногда бывает), изгоняют из общества, люди возвращаются в состояние нищеты. Это принято называть невезением. (C) Хайнлайн",

		"I can't give you a sure formula for success, but I can give you a formula for failure: try to please everybody all the time.",

		"Люди ни за что не стали бы терпеть такие высокие налоги, если бы не иллюзия, что богатые платят больше.",

		"Есть такая категория людей — они ссут против ветра и думают, что с ним борются. Вот также и в экономике: 'австрийцы' используют законы природы, а мэйнстрим — с ними борется.",

		"Capitalism is relatively new in human history. Prior to capitalism, the way people amassed great wealth was by looting, plundering, and enslaving their fellow man. Capitalism made it possible to become wealthy by serving your fellow man. (C) Walter E. Williams",

		"Stop trying to control. Let go of fixed plans and concepts, and the world will govern itself. The more prohibitions you have, the less virtuous people will be. The more weapons you have, the less secure people will be. The more subsidies you have, the less self-reliant people will be. (C) Lao Tzu",

		"Freedom is not worth having if it does not include the freedom to make mistakes. (C) Mahatma Gandhi",

		"Sometimes it is said that man cannot be trusted with the government of himself. Then how can he be trusted with the government of others? (C) Thomas Jefferson",

		"Our expensive welfare state is fueled by the destructive notion that 'greed' is when you want to keep your own money but 'compassion' is when you want to take somebody else's.",

		"It is irrational to make the claim that property is theft, considering that without property, there can be no theft at all. Either you respect possessions of the individual or you don't. You either believe an individual has a right to the fruits of his labor, or you don't.",

		"What kind of a society isn't structured on greed? The problem of social organization is how to deal with greed.",

		"There are people who think that plunder loses all its immorality as soon as it becomes legal. Personally, I cannot imagine a more alarming situation.",

		"People who accuse anarchists of being childish are confessing that they think everybody should be treated like children, that we need authorities to make us be good people. It's only on that premise that you can get to any comparison to 'childish' rebellion.",

		"We have to harm innocents in order to protect innocents from harm. That's the statists' logic.",

		"Socialism produces equal poverty.",

		"Liberals want the government to be your Mommy. Conservatives want government to be your Daddy. Libertarians want it to treat you like an adult.",

		"Every decent man is ashamed of the government he lives under.",

		"Paper money eventually returns to its intrinsic value, zero. (C) Voltaire",

		"'Need' now means wanting someone else's money. 'Greed' means wanting to keep your own. 'Compassion' is when a politician arranges the transfer. (C) Joseph Sobran ",

		"VOTE = Violating Others Through Extortion",

		"Democracy is the worship of jackals by jackasses. (C) H.L.Mencken",

		"Every election is a sort of advance auction sale of stolen goods. (C) H.L.Mencken",

		"Perhaps one of the most important accomplishments of my administration has been minding my own business. (C) Calvin Coolidge",

		"The urge to save humanity is almost always only a false-face for the urge to rule it.",

		"Equality of opportunity is freedom, but equality of outcome is repression.",

		"When buying and selling are controlled by legislation, the first things to be bought and sold are legislators.",

		"If you have been voting for politicians who promise to give you goodies at someone else's expense, then you have no right to complain when they take your money and give it to someone else, including themselves. (C) Thomas Sowell ",

		"Government is not reason; it is not eloquence; it is force. Like fire, it is a dangerous servant and a fearful master. (C) George Washington",

		"When under the pretext of fraternity, the legal code imposes mutual sacrifices on the citizens, human nature is not thereby abrogated. Everyone will then direct his efforts toward contributing little to, and taking much from, the common fund of sacrifices. Now, is it the most unfortunate who gains from this struggle? Certainly not, but rather the most influential and calculating. (C) Frederic Bastiat",

		"The plans differ; the planners are all alike. (C) Frederic Bastiat",

		"Attempting to control inflation through price controls is like attempting to control obesity by wearing tight clothing: the results are generally frustrating, often painful, and sometimes deeply embarrassing.",

		"The difference between libertarianism and socialism is that libertarians will tolerate the existence of a socialist community, but socialists can't tolerate a libertarian community.",

		"Society exists for the benefit of its members; not the members for the benefit of society. (C) Herbert Spencer",

		"Government is the great fiction through which everybody endeavors to live at the expense of everybody else.",

		"Inflation is taxation without legislation.",

		"A society that robs an individual of the product of his effort, is not strictly speaking a society, but a mob held together by institutionalized gang violence. (C) Ayn Rand",

		"The supply of government exceeds the demand by huge margin.",

		"I know no class of my fellowmen, however just, enlightened, and humane, which can be wisely and safely trusted absolutely with the liberties of any other class.",

		"It is amazing that people who think we cannot afford to pay for doctors, hospitals, and medication somehow think that we can afford to pay doctors, hospitals, medication and a government bureaucracy to administer it. (C) Thomas Sowell",

		"A free and prosperous society has no fear of anyone entering it. But a welfare state is scared to death of every poor person who tries to get in and every rich person who tries to get out.",

		"you can not multiply wealth by dividing it.",

		"Taxation of earnings from labor is on a par with forced labor. Seizing the results of someone's labor is equivalent to seizing hours from him and directing him to carry on various activities.",

		"Liberty is not a means to a higher political end. It is itself the highest political end.",

		"You can only be free if I am free.",

		"Have you ever noticed how statists are constantly 'reforming' their own handiwork? Education reform. Health-care reform. Welfare reform. Tax reform. The very fact that they're always busy 'reforming' is an implicit admission that they didn't get it right the first fifty times.",

		"Most protesters fight the system simply because they are not the ones in charge.",

		"The whole point of the system is to bribe you with the leftovers of what they have stolen from you in the first place.",

		"It is the greatest inequality to try to make unequal things equal. (C) Aristotle",

		"There are men in all ages who mean to govern well, but they mean to govern. They promise to be good masters, but they mean to be masters.",

		"Liberty is the only thing you cannot have unless you are willing to give it to others.",

		"The rule of the many by the few we call tyranny; the rule of the few by the many is tyranny also, only of a less intense kind. (C) Herbert Spencer ",

		"People who want the government to adopt and enforce their ideas are always the people whose ideas are idiotic.",

		"The principle that the majority have a right to rule the minority, practically resolves all government into a mere contest between two bodies of men, as to which of them shall be masters, and which of them slaves. A contest, that — however bloody — can never be finally closed, so long as man refuses to be a slave.",

		"When law and morality contradict each other, the citizen has the cruel alternative of either losing his moral sense or losing his respect for the law.",

		"Socialist system is so wonderful, that is should be forced upon people.",

		"We contend that for a nation to try to tax itself into prosperity is like a man standing in a bucket and trying to lift himself up by the handle. (C) Winston Churchill.",

		"Every new law adopted — another liberty lost.",

		"Ineptocracy — a system of government where the least capable to lead are elected by the least capable to produce, and where the members of society least likely to sustain themselves or succeed, are rewarded with goods and services paid for by the confiscated wealth of a diminishing number of producers.",

		"Let me offer you my definition of social justice: I keep what I earn and you keep what you earn. If you disagree, then tell me how much of what I earn belongs to you — and why? (C) Walter E. Williams",

		"The best argument against democracy is a five minute conversation with the average voter. (C) Chruchill",

		"То есть, безо всякого предварительного сговора, естественный ход событий, направляемый людьми, преследующими собственные интересы, приводит к результатам, которые не преследовал каждый чиновник в отдельности, но которые могут трактоваться, как выгодные системе в целом.",

		"Государства уже довольно давно используют один и тот же метод, с помощью которого рекламируют себя и свои услуги по решению проблем. Это метод состоит в поиске и даже создании этих проблем. Например, борьба с гомосексуализмом в России и у нас и борьба за привилегии гомосексуалистам на Западе — суть одно и то же, это один и тот же процесс, вызванный одними и теми же причинами. На самом деле, государству, как системе в целом, не важно за что и с кем бороться. Главное — бороться, иметь фронт работ и оправдание собственному существованию. Конкретное содержание борьбы в конкретной стране определяется улавливанием настроений.",

		"Politics is the art of looking for trouble, finding it whether it exists or not, diagnosing it incorrectly, and applying the wrong remedy.",

		"The only way to deal with an unfree world is to become so absolutely free that your very existence is an act of rebellion. (C) Albert Camus",

		"Socialism is a philosophy of failure, the creed of ignorance, and the gospel of envy. Its inherent virtue is the equal sharing of misery. (C) Winston Churchill",

		"A fine is a tax for doing wrong. A tax is a fine for doing well.",

		"Идеи левых не то чтобы плохие. Просто они не предназначены для людей в реальном мире. Поэтому левые стараются изменить человеческую природу. Будь то лагеря промывание мозгов в Китае и СССР или 'спецпрограммы' в американских школах. (C) Sowell",

		"Обычно 'социальные проблемы' происходят из факта, что у интеллектуалов есть теории, не подходящие реальному миру. И интеллектуалы уверены, что проблема в реальном мире, который надо менять. (C) Sowell",

		"Большинство крупных состояний заработано людьми, выросшими в куда худшей бедности, чем сегодняшние получатели пособий. (C) Sowell",

		"Борцы против богатства почему-то забывают, что богатство — это лучшее средство против бедности. (C) Sowell",

		"Государственные школы — это что-то особенное! В эпоху искусственного интеллекта они создают искусственную глупость.",

		"Почему те, кто изливает негодование на табачные компании, не особенно усердствуют в обличении производителей экипировки для экстремальных видов спорта? Смысл-то один и тот же: сознательный риск для получения удовольствия. (C) Sowell",

		"It's easy to be compassionate, when others are forsed to pay the cost.",

		"Демократия может сочетать в себе все виды гнета политического, религиозного, социального. Но, при демократическом строе, деспотизм становится неуловимым, так как он распыляется по различным учреждениям; он не воплощается ни в каком одном лице, он вездесущ и в то же время его нет нигде; оттого он, как пар, наполняющий пространство, невидим, но удушлив; он как бы сливается с национальным климатом. Он нас раздражает, от него страдают, на него жалуются, но не на кого обрушиться. Люди обыкновенно привыкают к этому злу и подчиняются. (C) Морис Палеолог",

		"Мечта социалистов — матриархальная модель государства, которое заботливо пеленает и кормит граждан, не особо интересуясь, хотят ли они спать или есть.",

		"Not since the days of slavery have there been so many people who feel entitled to what other people have produced as there are in the modern welfare state. (C) Sowell",

		"If you refuse to pay unjust taxes, your property will be confiscated. If you attempt to defend your property, you will be arrested. If you resist arrest, you will be clubbed. If you defend yourself against clubbing, you will be shot dead. These procedures are known as the rule of law.",

		"The smallest minority on earth is the individual. Those who deny individual rights cannot claim to be defenders of minorities.",

		"Легко защищать свободу слова, когда она применяется к правам тех, с кем мы согласны. И все же главная проверка связана с неоднозначными высказываниями — заявлениями, которые мы можем счесть дурными и отталкивающими.",

		"The original meaning of charity was not 'giving away someone else’s money'.",

		"The more one considers the matter, the clearer it becomes that redistribution is in effect far less a redistribution of free income from the richer to the poorer, as we imagined, than a redistribution of power from the individual to the State.",

		"If you are not free to choose wrongly and irresponsibly, you are not free at all.",

		"How does something immoral when done privately, become moral when it is done collectively?",

		"A democracy is nothing more than mob rule, where fifty-one percent of the people may take away the rights of the other forty-nine. (C) Thomas Jefferson",

		"Capitalism is private ownership of property, socialism is state ownership of people.",

		"If a businessman makes a mistake, he suffers the consequences. If a bureaucrat makes a mistake, you suffer the consequences. (C) Ayn Rand",

		"Социалисты считают, что получать прибыль — грех. Я же считаю, что настоящий грех — это терпеть убытки. (C) Уинстон Черчиль",

		"Socialists believe that profit is a sin. I believe that the real sin is to suffer losses. (C) Churchill",

		"No man is good enough to govern the other without the other man's consent. (C) Lincoln",

		"Лучше стороить систему на существующих пороках, чем на несуществующих добродетелях, как это делают социалисты.",

		"Пять человек находится в комнате. По причине, что трое из них разделяют одно мнение, а двое других — имеют противоположное, трое имеют моральное право принуждать двоих? Такое впечатление, будто они получают магическую силу только потому, что их больше.",

		"До тех пор, пока их было два на два - каждый оставался сам себе хозяин. Но в тот момент, когда один человек присоединился к первой либо второй группе, тут же члены этой группы стали владельцами меньшинства.",

		"Была ли когда-либо настолько примитивная и безапелляционная власть большинства?",

		"Правительство как грудной младенец: чудовищный аппетит на одном конце и полная безответственность на другом.",

		"We disapprove of state education. Then the socialists say that we are opposed to any education. We object to a state religion. Then the socialists say that we want no religion at all. We object to a state-enforced equality. Then they say that we are against equality. (C) Bastiat",

		"People would rather be equal in slavery than unequal in freedom.",

		"Если вы хотите получить подтверждение того, что демократия не выполняет своих обещаний, обратите внимание на то, что политики перед каждыми выборами признают, что правительство все испортило. Каждый раз они обещают все изменить — образование, безопасность, здравоохранение и т.д. — к лучшему. Но они часто предлагают то же самое решение: дать им больше денег и власти, что бы они решили эти проблемы. И конечно, таким образом проблемы никогда не решаются, так как они вызваны деньгами и властью тех же самых политиков.",

		"When we give government the power to make medical decisions for us, we, in essence, accept that the state owns our bodies.",

		"Three questions that would destroy most of the arguments on the left: Compared to what? At what cost? What hard evidence do you have? (C) Sowell",

		"One of the great mistakes is to judge policies and programs by their intentions rather than their results.",

		"Mercy to the guilty is cruelty to the innocent.",

		"Of all tyrannies, a tyranny sincerely exercised for the good of its victims may be the most oppressive. It would be better to live under robber barons than under omnipotent moral busybodies. The robber baron's cruelty may sometimes sleep, his cupidity may at some point be satiated; but those who torment us for our own good will torment us without end for they do so with the approval of their own conscience. (C) C. S. Lewis",

		"Socialism is a great system until you run out of other people's money.",

		"Unsatisfied demand is not a defect of a planned economy. On the contrary — that is its only advantage! In a market economy supply is determined by demand, whereas in a planned economy supply is determined by the plan. Planned-economy proponents believe that a person's wants do not correspond to his real needs, so common resources are expended excessively in some areas and less than necessary in others. Therefore, a bureaucrat is needed to determine for each person his needs and plan for their satisfaction. So, if you think your needs have not been satisfied, you're mistaken. It's not up to you to decide.",

		"I prefer dangerous freedom over peaceful slavery.",

		"Вера в государство — это вера в то, что человек, распоряжающийся вашими деньгами, тратит их настолько эффективнее вас, что окупает этим расходы на свое содержание.",

		"If you believe in equal rights, then what do 'women's rights' or 'gay rights' even mean? Either they are redundant or they are violations of the principle of equal rights for all.",

		"The most dangerous thing about marijuana is getting caught with it, and it's inflicted by government.",

		"Мечта рабов: рынок, где можно было бы покупать себе хозяев. (C) Лец",

		"When resisting a criminal, you obstruct investigation of your murder!",

		"Демократия — это когда власти уже не назначаются безнравственным меньшинством, а выбираются безграмотным большинством. (C) Бернард Шоу",

		"Я всё никак не мог понять, зачем чиновники запретили наркотические вещества. Мне не был ясен смысл этого запрета: он не работает, качество наркотических веществ становится хуже, люди продолжают погибать. Но ответ лежит на поверхности. Целью запрета было создание проблемы, с которой можно вечно бороться.",

		"The public school system is a twelve year sentence of mind control. Crushing creativity, smashing individualism, encouraging collectivism and compromise, destroying the exercise of intellectual inquiry, twisting it instead into meek subservience to authority.",

		"Отечество раба там, где палка. (C) Генрих Гейне",

		"Свобода выбора позволяет людям делать ошибки, однако отсутствие свободы намного хуже.",

		"Позвольте мне предложить вам мое определение социальной справедливости: я владею тем, что я зарабатываю, и вы владеете тем, что вы зарабатываете. Если вы не согласны, то скажите, какая часть моей зарплаты принадлежит вам, и на каком основании?",

		"Всякий, кто борется с коррупцией иначе, чем ограничением полномочий чиновника, этим увеличивает влияние государства на нашу жизнь, и соответственно возможности для произвола чиновников. Если он понимает это, то он фашист; если нет — он дурак.",

		"Истинная цель государства заключается в том, чтобы люди свободно владели своей собственностью и чтобы они не подвергались опасности. (C) Цицерон",

		"Вещь стоит столько, сколько за нее можно взять с покупателя. (C) Цицерон",

		"Рабы всех страстей сердятся на чужие пороки так, словно им завидуют, и тяжелее всего наказывают тех, кому больше всего им хотелось бы подражать. (C) Цицерон",

		"The emperor had licensed many so-called monopolies, selling the welfare of his subjects to those who did not shun from this abomination, after he had exacted his price for the privilege. (C) Procopius, VI century AD",

		"Одним из следствий 'права на социальные льготы' является то, что люди, не дающие ничего обществу, почему-то считают, что общество должно им нечто, — уж не за то ли, что они почтили нас своим присутствием? (C) Томас Соуэлл",

		"Бог сделал право на труд достоянием каждого человека; и это достояние есть первое, самое священное и неотъемлемое из всех. (C) Людовик XVI, Эдикт об Упразднении Гильдий",

		"It would be better that England should be free than that England should be compulsorily sober. (C) William Connor Magee, archbishop of York, on prohibition",

		"Есть нечто более жестокое, чем личная жестокость. Это холодная жестокость бездушной системы.",

		"Частного человека, который с оружием в руках отнимет у соседа корову или десятину посева, сейчас же возьмут полицейские и посадят в тюрьму. Кроме того, такого человека осудит общественное мнение, его назовут вором и грабителем. Совсем иное с государствами: все они вооружены, власти над ними нет никакой. И главное то, что общественное мнение, которое карает всякое насилие частного человека, восхваляет, возводит в добродетель патриотизма всякое присвоение чужого для увеличения могущества своего отечества. (C) Л.Н.Толстой",

		"Liberty is not collective, it is personal. All liberty is individual liberty... When once the right of the individual to liberty and equality (before the law) is admitted, there is no escape from the conclusion that he alone is entitled to the rewards of his own industry. Any other conclusion would necessarily imply either privilege or servitude. (C) Calvin Coolidge",

		"Когда государство начинает убивать, оно всегда называет себя Родиной.",

		"The more numerous the laws, the more corrupt the government. (C) Tacitus",

		"The real division is not between conservatives and revolutionaries but between authoritarians and libertarians. (C) George Orwell",

		"Мы избегали методов, свойственных социальному государству, потому что видели, как великий британский народ в результате социалистической уравниловки превратился в посредственный. (C) Ли Кван Ю",

		"If you look at the drug war from a purely economic point of view, the role of the government is to protect the drug cartel from competition. (C) Milton Friedman",

		"Если в стране платят за бедность, болезни и старость, то население через некоторое время становится бедным, больным и старым.",

		"Прогрессивная общественность никак не может взять в толк, что коррупция существует только потому, что она ДЕШЕВЛЕ государственного регулирования. И отсюда есть два вывода. Первый: не нравится коррупция — сокращай регулирование. Второй: если заставить всех 'делать по закону', то не будет вам товаров и услуг.",

		"Механитис – профессиональное заболевание тех, кто верит, что ответ математической задачи, которую он не может ни решить, ни даже сформулировать, легко будет найти, если получить доступ к достаточно дорогой вычислительной машине. (C) из книги 'Физики Шутят'",

		"Назвав инфляцией рост цен, а не рост денежной массы, государство канализировало недовольство людей против буржуев, торговцев и прочих своих естественных врагов.",

		"Если ты за свободу слова, но запрещаешь высказывания, которые тебе не нравятся, то ты против свободы слова. Это такая элементарная проверка на уровне формальной логики.",

		"The 'private sector' of the economy is, in fact, the voluntary sector; and the 'public sector' is, in fact, the coercive sector. (C) Henry Hazlitt",

		"A license is what you get when the government steals your rights and then sells them back to you in the form of temporary privileges.",

		"Любовь к свободе — это любовь к людям; любовь к власти — это себялюбие. (C) Хэзлитт",

		"Jailing someone for tax evasion — is like jailing the victim of mugging for not also giving their wrist watch along with their wallet.",

		"Главное зло государственного устройства не в уничтожении жизней, а в уничтожении любви и возбуждении разъединения между людьми. (C) Лев Толстой",

		"Socialism — the ridiculous notion that people are better off when you take their stuff, limit their choices, make them dependent upon politicians and concentrate power.",

		"Der Staat will nicht bloss, dass Du gehorchst. Der Staat will, dass Du gehorchen willst. (C)  Mencken",

		"Libertarianism is the radical notion that other people are not your property.",

		"Когда обезьяна взяла в руки палку, она стала человеком. А когда она начала совать её в колёса, то стала чиновником.",

		"Only in America could the 'rich people', who pay 86%% of all income taxes, be accused of not paying their 'fair share' by people who don't pay any income taxes at all.",

		"Бесплодность полицейских мер обнаруживала всегдашний прием плохих правительств — пресекая следствия зла, усиливать его причины. (C) Ключевский",

		"Чем выше спрос, тем ниже цена, которую нужно платить за свободу. (C) Станислав Лец",

		"Whenever I pass by a school or a jail, I always feel sorry for the people inside.",

		"The system is a retarded giant throwing wads of money and books of rules in random directions while shouting “LOOK AT ME! I’M HELPING! I’M HELPING!” Sometimes by luck you catch a wad of cash, and you think the system loves you. Other times by misfortune you get hit with a book, and you think the system hates you. But either one is giving the system too much credit.",

		"Disregard for the preferences and interests of individuals alive today in order to pursue some distant social goal that their rulers have claimed is their duty to promote has been a common cause of misery for people throughout the ages. (C) Isaiah Berlin",

		"An unjust law is no law at all. (C) St. Augustine",

		"Justice being taken away, then, what are kingdoms but great robberies? For what are robberies themselves, but little kingdoms? (C) St. Augustine",

		"Statism is the utopian ideal that just the right amount of violence used by just the right people in just the right way can perfect society.",

		"Corporations can buy favors from government only because government has favors to sell.",

		"Laws are opinions with guns.",

		"In a democracy people lose thier rights as soon as they are outnumbered.",

		"Никто не повинен в том, если он родился рабом; но раб, который не только чуждается стремлений к своей свободе, но оправдывает и прикрашивает свое рабство, такой раб есть вызывающий законное чувство негодования, презрения и омерзения холуй и хам. (C) Ленин",

		"It is difficult to free fools from the chains they revere. (C) Voltaire",

		"A demagogue is one who preaches doctrines he knows to be false to people he knows to be idiots. (C) Mencken",

		"Communism is a system where everything is supposed to be free, except you.",

		"Government is like a portable toilet. I'd rather not use it, but I'll be arrested if I do my business outside of it.",

		"В западных странах люди никогда не стали бы терпеть такой высокий уровень налогов, если бы не ИЛЛЮЗИЯ что богатые платят больше. И вот правительства вынуждены одной рукой продавать народу эту иллюзию, а другой делать послабления бизнесу, иначе экономика просто развлится. И пока пипл хавает, недовольство плебса направлено на бизнес, а не на правительство. Правительство довольно, бизнес вынужден подыгрывать, а плебеи сами виноваты, что голосуют за этот лохотрон. ",

		"Придумав название 'золотой век Коммода', император распорядился снизить цены, чем вызвал затем еще большую нужду. (C) Элий Лампридий",

		"The law, in its majestic equality, commands rich and poor alike to comply with KYC, AML, and megabytes of other regulations.",

		"The poor do not so much break the law as the law breaks them. When laws and regulations reach sufficient complexity and burden, they serve as a barrier to competition and mean that only the rich can navigate them. Everyone else becomes accidental criminals.",

		"This is why large banks lobby for more regulation. It’s costly to the big banks, but they can afford the armies of lawyers and lobbyists, whereas small community banks can’t afford compliance and succumb to a cheap buyout.",

		"There’s also a political element. A tyrant doesn’t want a nation of law abiding citizens — those are impossible to intimidate. A nation with laws so numerous and complex that everyone violates some grants the tyrant the power of selective enforcement.",

		"Идеальный человек тоталитарного режима — не убежденный нацист или коммунист, а тот, для кого различие между фактом и вымыслом, правдой и ложью больше не существует.",

		"Тому, кто не видит разницы между властью денег и властью плётки, следовало бы на своей шкуре испробовать одно и другое.",

		"For stupid people the success of others is a grievance, rather than an example, even when they enjoy the fruits of this success.",

		"I would rather have questions that can't be answered than answers that can't be questioned. (C) Richard Feynman",

		"It's very easy to scare people and start a political stampede. There are people who could be upset if they were told that half of all Americans earn less than the median income. (C) Sowell",

		"Socialist wants everything you have except your job.",

		"Banning weapons is the most white privilege idea ever. Only rich liberals scoff at the notion that a man might need to defend his own life. It's the personal security equivalent of 'just have a maid to do it for you!'",

		"I want to carry a gun because policeman is too heavy.",

		"Наша система зиждется на невысказанном предположении, что чиновник расходует чужие деньги настолько эффективнее их хозяина, что даже окупает этим своё содержание.",

		"Capitalism turns luxuries into necessities. Socialism turns necessities into luxuries.",

		"Socialism in general has a record of failure so blatant that only an intellectual could ignore or evade it. (C) Sowell",

		"There is no point blaming the tragedies of socialism on the flaws or corruption of particular leaders. Any system which allows some people to exercise unbridled power over others is an open invitation to abuse, whether that system is called slavery or socialism or something else. (C) Sowell",

		"A statement may be both true and dangerous. The previous sentence is such a statement. (C) David Friedman",

		"A writer who says that there is no truth, or that all truth is 'merely relative', is asking you not to believe him. So don't.",

		"We should not be surprised to find the left concentrated in institutions where ideas do not have to work in order to survive.",

		"Определение левизны: стремление творить добро за чужой счет.",

		"Egalitarians create the most dangerous inequality of all - inequality of power. Allowing politicians to determine what all other human beings will be allowed to earn is one of the most reckless gambles imaginable.",

		"В этатистской картине мира есть два типа субъектов. Первый - это добрые и неподкупные ангелы (или роботы) - чиновники и политики. Второй - злые и своекорыстные людишки, о которых первый тип заботится, отдавая им беспристрастные и направленные только на их благо приказы. Для этатиста последствия, в конечном итоге, определяются (или должны определяться) приказами добрых роботов. Если добрые роботы сказали, что минимальная зарплата будет 25 долларов в час, то единственным последствием этого должно стать увеличение благосостояния трудящихся. Если этого не происходит, нужно просто найти робота, который поломался, и заменить его.",

		"No matter what anyone may say about making the rich and the corporations pay taxes, in the end they come out of the people who toil. (C) Calvin Coolidge",

		"For those looking for security, be forewarned that there's nothing more insecure than a political promise. (C) Harry Browne",

		"Я не участвую в выборах. В самом деле, не стану же я добровольно голосовать за человека, который стремится получить власть надо мной.",

		"I considered selling my weapons 'back' to the government, but after a background check and thorough investigation into the buyer, I determined the buyer has a history of violence and is mentally unstable.",

		"The state is mafia pretending to be a human rights organization.",

		"Whatever the issue, let freedom offer us a hundred choices, instead of having government force one answer on everyone. (C) Harry Browne",

		"Общественное благо лучше всего достигается в процессе преследования людьми собственных частных интересов. Это не дает покоя тем, кого больше волнуют добрые намерения, чем результат.",

		"I don't think the key aspect of capitalism is 'selfishness'. The key aspect is 'voluntary'. That allows both altruism and selfishness to exist.",

		"Democracy and liberty are not the same. Democracy is little more than mob rule, while liberty refers to the sovereignty of the individual.",

		"Without anarchy, there would be chaos. (C) Jeffrey Tucker",

		"Public schools are literal prisons for children, and the only place many people will ever encounter physical violence in their lives. Or, it's the only place where people can phisically assault weaker victims with impunity.",

		"Во всякой стране капитализм присутствует ровно настолько, насколько там вообще что-то производится и функционирует. Коллективистские системы декларируют уничтожение капитализма, но на самом деле они просто паразитируют на нём.",

		"What makes the difference between a gang and a state is the belief that there is a difference between a gang and a state. (С) Jakub Bożydar Wiśniewski",

		"Socialists insist that free enterprise capitalism 'puts profits above people'. Nonsense! Profits go to those who most efficiently serve the needs of people.",

		"Without profit, we would be living hand to mouth. Profit is what allows us to advance our standard of living.",

		"Profit is also the one and only reliable incentive for efficiency. It literally measures efficiency. Which is why, when try to remove the profit motive, efficiency collapses.",

		"It’s actually unbelievable that it’s super common for people to have school nightmares into adulthood and we as a society just accept that level of trauma as normal.",

		"The secret to happiness is freedom. And the secret to freedom is courage.",

		"When one gets in bed with government, one must expect the diseases it spreads. (C) Ron Paul",

		"The average man's love of liberty is nine-tenths imaginary, just like his love of sense, justice and truth... Liberty is not a thing for the great masses of men. It is the exclusive possession of a small and disreputable minority, like knowledge, courage and honor. It takes a special sort of man to understand and enjoy liberty, and he is usually an outlaw in democratic societies. (C) Mencken",

		"Mankind will soon discover that unbridled majorities are as tyrannical and cruel as unlimited despots. (C) John Adams",

		"The whole idea of our government is this: if enough people get together and act in concert, they can take something and not pay for it. (C) P.J. O'Rourke",

		"...не только не видят своего рабства, но гордятся им, считая себя свободными гражданами великих государств, гордятся этим так же, как лакеи гордятся важностью господ, которым они служат. (C) Лев Толстой",

		"Утопическими следует признать не анархические убеждения, а наивные верования, что рвущиеся к власти социопаты желают тебе помочь.",

		"Since slavery has been discredited, 'democracy' is the only remaining rationale for state injustice to be tolerated.",

		"Конструктору полагается стоять под мостом, когда по нему поедет первый поезд. Капитану полагается последним уходить с корабля. И только чиновник норовит свалить подальше от того места, где он улучшил жизнь по новой прогрессивной методе.",

		"В промышленности люди пользуются технологиями космического века, а в управлении государством - технологиями бронзового века.",

		"Если государство оставит гражданам хоть немого свободы, всегда найдутся холуи, которые будут жаловаться на 'дыры в законах'.",

		"No matter how much the government controls the economic system, any problem will be blamed on whatever small zone of freedom that remains.",

		"The state is a vast enterprise for declaring all sorts of things legal for itself and illegal for us. (C) Lew Rockwell",

		"We are living in an era of 'woke' capitalism in which companies pretend to care about social justice to sell products to people who pretend to hate capitalism.",

		"The individual almost always pursues their personal interests over the needs of wider society, therefore if you want your civilization to thrive, you must be sure to align the incentives of the individual with the needs of the collective.",

		"Тот, кому чужда жизнь, кто неспособен к ней, тому ничего больше не остается, как стать чиновником. (C) Антон Чехов",

		"If you cannot achieve equality of performance among people born to the same parents and raised under the same roof, how realistic is it to expect to achieve it across broader and deeper social divisions? (C) Sowell",

		"Борцы за социальную справедливость во всех бедах нашего мира винят капиталистов, т.е. продавцов товаров и услуг, тогда как покупатели у них якобы сильно страдают от добровольных транзакций. Но когда доходит до секса, то вдруг вся их логика выворачивается наизнанку. Если женщина продаёт своё тело, то виноваты уже покупатели, а продавцы превращаются в невинных жертв.",

		"There's a simple way to keep money out of politics: Keep politics out of our money.",

		"Taxation of earnings from labor is on a par with forced labor. Seizing the results of someone's labor is equivalent to seizing hours from him and directing him to carry on various activities. (C) Robert Nozick",

		"'Когда среди 100 человек один властвует над 99 - это несправедливо, это деспотизм; когда 10 властвуют над 90 - это также несправедливо, это олигархия; когда же 51 властвует над 49 - тогда это совершенно справедливо, это свобода!' Может ли быть что-нибудь смешнее такого рассуждения? А между тем, это самое рассуждение служит основой деятельности всех улучшателей государственного устройства. (С) Лев Толстой",

		"Free men are not equal and equal men are not free. (C) Lawrence Reed",

		"Arguing that you don't care about privacy because you have nothing to hide is no different than saying you don't care about free speech because you have nothing to say.",

		"If you think that privacy is for criminals, then lead by example - publish all your emails, texts, bank statements, your credit card PIN, and maybe a live stream from your bathroom.",

		"Populism is what we call democracy when we don't like the outcome.",

		"I believe that all government is evil, and that trying to improve it is largely a waste of time. (C) H. L. Mencken",

		"KYC is an affront to basic dignity. The idea that adults must be supervised with their funds, require permission to transact and must submit to judgment on how they use their money - contradicts the concept of private property.",

		"Taxation is thought to be indispensable to civilization today, just as slavery once was. Advocates of taxation claim that since most people pay assigned taxes before the guns show up, they have implicitly agreed to it as the price of living in 'society'. Most slaves obeyed their master before he got out the whip, yet we would hardly argue that this constituted agreement to their servitude. Today, we have an enlightened perspective on slavery, just as one day we will have an enlightened perspective on taxes and other forms of aggression we now think of as 'the only way'.",

		"Property rights are understood by everyone except babies and socialists. A lot of animals understand them too.",

		"When does plunder stop? It stops when it becomes more painful and more dangerous than labor. (C) Bastiat",

		"Thou shalt not steal, even by majority vote.",

		"Give the government an inch and it will take a mile. But it will take it inch by inch.",

		"Some people ask the socialists: How much of your neighbor's income belongs to you and why? But I would rather ask: How much of your income belongs to me and why?",

		"There is no such thing as a fair share of taxes. The fair price of a good or service is that which both parties voluntarily agree to. But taxes are collected by force. What is the 'fair' rate of robbery? For robbers - whatever they can get away with.",

		"The most important single central fact about a free market is that no exchange takes place unless both parties benefit. (C) Milton Friedman",

		"No government knows any limits to its power except the endurance of the people. (C) Lysander Spooner",

		"The difference between libertarianism and socialism is that libertarians will tolerate the existence of a socialist community, but socialists can't tolerate a libertarian community.",

		"Render unto Caesar what is Caesar's, and outside of Caesar's realm use Bitcoin. (C) Satoshi Nakamoto",

		"Wealth creation is an evolutionarily recent positive-sum game. Status is an old zero-sum game. Those attacking wealth creation are often just seeking status.",

		"The only people, who actually 'hate you for your freedom', are your own politicians.",

		"Власть рождает паразитов, да здравствует анархия! (C) Нестор Махно",

		"You are being conditioned to view your freedom as selfish. Stop feeling guilty for doing what's best for you.",

		"The same people that demand free services would be outraged if they were forced to provide those services for free.",

		"Burocrats can stay retarded longer than we can stay solvent.",

		"In Bitcoin you can not exclude even the least powerful man in the world.",

		"Everything the State says is a lie, and everything it has it has stolen. (C) Friedrich Nietzsche",

		"In individuals insanity is rare. But in groups, parties, nations and epochs - it is the rule. (C) Friedrich Nietzsche",

		"What is a loophole? If the law does not punish a definite action or does not tax a definite thing, this is not a loophole. It is simply the law. (C) Ludwig von Mises",

		"The states are responsible for the deaths of hundreds of millions of people and immeasurable destruction in the 20th century alone. Compared to that, the victims of private crimes are almost negligible.",

		"Socialism is the belief that slavery is OK, as long as the State does it rather than private individuals.",

		"Any sufficiently complicated legal system is indistinguishable from saying 'LOL, fuck you' to all the peasants who can't afford expensive lawyers. (C) Eliezer Yudkowsky",

		"Зачем громкие слова против меритократии? Просто на деле начинай поддерживать тупых и рукожопых - иди к хреновому врачу, покупай товары плохого качества, смотри скучные фильмы, читай нудные книги, и т.д.",

		"Democracy is great, except when our rivals win. Then we call it populism.",

		"If government is so good, why do we have so few of those? We should have millions of governments, instead of 200.",

		"If you voted, you have no right to complain. You wanted a master, and you got one.",

		"Fiat is kleptocurrency because whoever in charge is unable to resist the urge to steal.",

		"If you listen to modern marxists, it turns out that it's not capitalists who oppress workers, but workers who oppress welfare recipients suffering from obesity.",

		"Some people want independence, other people want power. Whom you will support?",

		"In any government bureaucracy, they are not working for you, but for the mythical 'public sector', which is really just a stash of stolen money divided among the robber class. (C) Lew Rockwell",

		"The idea of social justice is that the state should treat different people unequally in order to make them equal. (C) Friedrich Hayek",

		"It is not the fault of capitalism that the common man does not appreciate uncommon books. (C) Ludwig von Mises",

		"The only things government can do are regulate and redistribute, prohibit and penalize, confiscate and command. Are these the things that liberty is made of? Somebody else's money and an endless list of Thou Shalt Nots? (C) James Bovard",

		"Freedom is never more than one generation away from extinction. (C) Ronald Reagan",

		"A group of people who think differently is a market. A group of people who think alike is a mob.",

		"Консерватор это коллективист, которому хотят причинить добро другие коллективисты, а он по справедливости хотел бы, чтобы всё было наоборот.",

		"Some people believe that inequality based on race or sex is wrong, but inequality based on productivity is right.",

		"The first lesson of economics is scarcity: there is never enough of anything to fully satisfy all those who want it. The first lesson of politics is to disregard the first lesson of economics. (C) Thomas Sowell",

		"Your enemies' rights are a canary for your own rights.",

		"Для всех государственных сервисов внезапный наплыв клиентов - не подарок судьбы, а досадная помеха.",

		"The United States is a nation of laws, badly written and randomly enforced. (C) Frank Zappa",

		"Greedy capitalists get money by trade. Compassionate liberals steal it. (C) David Friedman",

		"Easiest litmus test for whether someone's ideas are garbage: do they think the government should enforce their ideas? or do they intend to convince people peacefully?",

		"Ask not what your government can do for you. Ask what your government can do to you.",

		"И сотворили школу так, как повелел им дьявол. Ребенок любит природу, поэтому его замкнули в четырех стенах. Он не может сидеть без движения — его принудили к неподвижности. Он любит работать руками, а его стали обучать теориям и идеям. Он любит говорить — ему приказали молчать. Он стремится понять — ему велели учить наизусть. Он хотел бы сам искать знания — ему их дают в готовом виде. И тогда дети научились тому, чему никогда бы не научились в других условиях. Они научились лгать и притворяться.",

		"If you allow the government to break the law during emergencies, it will create emergencies to break the law.",

		"It is self-evident that no number of men, by conspiring, and calling themselves a government, can acquire any rights whatever over other men, or other men's property, which they had not before, as individuals. (C) Lysander Spooner",

		"Democracy: you say what you like and do what you are told. ",

		"When you take the free will out of education, that makes it schooling. (C) John Taylor Gatto",

		"The only idea they have ever manifested as to what is a government of consent: everybody must consent under threat of jail. (C) Lysander Spooner",

		"The bureaucracy is expanding to meet the needs of the expanding bureaucracy.",

		"I would rather have questions that can't be answered than answers that can't be questioned. (C) Richard Feynman",

		"Free men are not equal and equal men are not free.",

		"If you have to choose one of the seven deadly sins, for goodness sake, don't pick envy. It is the only sin that gives you no pleasure whatsoever, not even fake or temporary satisfaction.",

		"Most of the trouble in this world has been caused by folks who can't mind their own business, because they have no business of their own to mind.",

		"The closer people live to the consequences of their beliefs the more rational they become. People would slap their hands with a hammer way more often, if pain was delayed by a year.",

		"Государство пресекает всякую несправедливость, кроме той, которую оно творит само. (С) Ибн Хальдун",

		"Luckily, we do not get as much government as we pay for.",

		"Is there really someone who, searching for a group of wise and sensitive persons to regulate him for his own good, would choose that group of people that constitute the membership of both houses of Congress? (C) Robert Nozick",

		"Those who promise us paradise on earth never produced anything but a hell. (C) Karl Popper",

		"Being right too soon is socially unacceptable. (C) Robert A. Heinlein",

		"If a system constantly produces a different outcome than the one it is 'intended' for then it's perfectly reasonable to assume the actual intention is the outcome it continues to produce.",

		"Some people are only fit to be slaves. But I reject slavery because no men are fit to be masters.",

		"Вместо того, чтобы честно сказать: 'Да, я подлец и скотина', люди пускаются в какие-то нудные рассуждения о том, что на фоне общего благополучия, причинённый ихними регуляциями ущерб незначителен.",

		"Люди не равны, и от отрицания этого факта равнее не становятся.",

		"A duty to others undertaken voluntarily is divine. A duty to others imposed by force is slavery.",
	}
	return s
}
