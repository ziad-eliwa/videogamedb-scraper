USE VideoGames;

LOAD DATA LOCAL INFILE '/home/ziad-eliwa/mobygames-scraper/csv/Games.csv' IGNORE
INTO TABLE Game
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(Name, CoverURL, Description);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/People.csv' IGNORE
INTO TABLE People
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(MobyID, Name);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/gamedirector.csv' IGNORE
INTO TABLE GameDirector
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(GameName, DirectorID, DirectorRole);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/platform.csv' IGNORE
INTO TABLE Platforms
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(Name, Description, Years_Active);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/category.csv' IGNORE
INTO TABLE Category
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(CategoryName, CategoryType, Description);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/GameCategory.csv' IGNORE
INTO TABLE GameCategory
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(GameName, CategoryName, CategoryType);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/companies.csv' IGNORE
INTO TABLE Company
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(Name, Overview, LogoURL, YearStarted, YearEnded, Country);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/GameRelease.csv' IGNORE
INTO TABLE GameRelease
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 LINES
(@GameName, @PlatformName, @ReleaseDate, @ReleaseComment, @RetailPrice)
SET
GameName = @GameName,
PlatformName = @PlatformName,
ReleaseDate = STR_TO_DATE(@ReleaseDate, '%M %d, %Y'),
ReleaseComment = @ReleaseComment,
RetailPrice = @RetailPrice;

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/Developer.csv' IGNORE
INTO TABLE Developer
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(DeveloperName, GameName);


-- Load Publisher data
LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/Publisher.csv' IGNORE
INTO TABLE Publisher
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(PublisherName, GameName);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/tech-specs.csv' IGNORE
INTO TABLE TechSpecs
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(SpecName, SpecAttribute);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/TechSpecsRelease.csv' IGNORE
INTO TABLE TechSpecRelease
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(GameName, PlatformName, SpecName, SpecAttribute);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/maturity-rating.csv' IGNORE
INTO TABLE MaturityRating
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(MaturityRatingType, MaturityRatingAttribute);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/MaturityRatingRelease.csv' IGNORE
INTO TABLE MaturityRatingRelease
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(GameName, PlatformName, MaturityRatingType, MaturityRatingAttribute);

LOAD DATA LOCAL INFILE '~/mobygames-scraper/csv/MobyGameReview.csv' IGNORE
INTO TABLE MobyGameReview
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(GameName, PlatformName, MobyScorePlayers, MobyScoreCritic, NumberOfPlayers, NumberOfCritics, MobyRating);