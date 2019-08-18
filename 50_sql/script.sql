use LearnGolang;
create table if not exists `tb_student`(
                                           `id` nvarchar(5) not null,
                                           `name` nvarchar(100) not null,
                                           `age` int(11) not null,
                                           `grade` int(11) not null
) engine=InnoDB default charset = latin1;


INSERT INTO `tb_student` (`id`, `name`, `age`, `grade`) VALUES
('B001', 'Jason Bourne', 29, 1),
('B002', 'James Bond', 27, 1),
('E001', 'Ethan Hunt', 27, 2),
('W001', 'John Wick', 28, 2);

ALTER TABLE `tb_student` ADD PRIMARY KEY (`id`);
