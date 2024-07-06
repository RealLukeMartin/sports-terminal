"""The main entry point for the data scraper."""
import espn_scraper as espn

leagues = espn.get_leagues()

print(leagues)
