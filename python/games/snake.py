""" Imports """
import pygame as pg
from random import randrange
import csv
from datetime import datetime
import os


""" Functions """
def get_food_position(food):
    food.center = get_random_position()
    while food.center in [body.center for body in snake_body]:
        food.center = get_random_position()
    return food.center

def show_menu(screen):
    # Set the font and font size
    # pg.font.init()
    font = pg.font.Font(None, 36)

    # Set up the menu options
    menu_options = ['Easy', 'Medium', 'Hard', 'Scoreboard', 'Quit']
    menu_length = len(menu_options)

    # Set up the menu cursor
    cursor_pos = 0

    # Loop until the user selects an option
    while True:
        # Handle events
        for event in pg.event.get():
            if event.type == pg.QUIT:
                exit()
            elif event.type == pg.KEYDOWN:
                if event.key == pg.K_UP:
                    cursor_pos = (cursor_pos - 1) % menu_length
                elif event.key == pg.K_DOWN:
                    cursor_pos = (cursor_pos + 1) % menu_length
                elif event.key == pg.K_RETURN:
                    return menu_options[cursor_pos]

        # Draw the menu options
        screen.fill('black')

        for i, option in enumerate(menu_options):
            if i == cursor_pos:
                color = 'red'
            else:
                color = 'white'
            text = font.render(option, True, color)
            text_rect = text.get_rect(center=(screen.get_width()/2, 100 + i*50))
            screen.blit(text, text_rect)

        # Update the screen
        pg.display.flip()

def insert_scoreboard(score):
    if score == 0:
        return
    scoreboard = get_scoreboard() if get_scoreboard() else []
    date = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    scoreboard.append([date, score])
    with open('/tmp/.snake.scoreboard', 'w', newline='') as file:
        writer = csv.writer(file)
        for row in scoreboard:
            writer.writerow(row)

def get_scoreboard():
    try:
        with open('/tmp/.snake.scoreboard', 'r') as file:
            reader = csv.reader(file)
            return list(reader)
    except FileNotFoundError:
        with open('/tmp/.snake.scoreboard', 'w', newline='') as file:
            pass


""" Screen Configurations """
window = 400
tile_size = 15
range = (tile_size // 2 + 20, window - tile_size // 2, tile_size)
screen = pg.display.set_mode((window, window))
pg.display.set_caption('Snake')

clock = pg.time.Clock()
pg.font.init()

""" Game Configurations """
get_random_position = lambda: (randrange(*range), randrange(*range))
time, snake_speed = 0, 0
start_screen = True
mode = {
    'easy': {
        'snake_speed': 150,
        'infinite': True,
    },
    'medium': {
        'snake_speed': 110,
        'infinite': False,
    },
    'hard': {
        'snake_speed': 70,
        'infinite': False,
    },
}
snake_pos = get_random_position()
snake = pg.rect.Rect([0, 0, tile_size - 2, tile_size - 2])
snake.center = get_random_position()
length = 1
snake_body = [snake.copy()]
snake_color = 'white'
snake_dir = (0, 0)
food = snake.copy()
food.center = get_food_position(food)
food_color = 'red'

""" Game Over Configurations """
score = 0
game_over = False
self_collision = False
score_font = pg.font.SysFont('arialunicode', 20)
score_color = 'white'
game_over_font = pg.font.SysFont('arialunicode', 40)
game_over_color = 'red'

""" Main Loop """
while not game_over:
    for event in pg.event.get():
        if event.type == pg.QUIT:
            game_over = True
        if event.type == pg.KEYDOWN:
            if (event.key == pg.K_UP or event.key == pg.K_w) and snake_dir != (0, tile_size):
                snake_dir = (0, -tile_size)
            if (event.key == pg.K_DOWN or event.key == pg.K_s) and snake_dir != (0, -tile_size):
                snake_dir = (0, tile_size)
            if (event.key == pg.K_LEFT or event.key == pg.K_a) and snake_dir != (tile_size, 0):
                snake_dir = (-tile_size, 0)
            if (event.key == pg.K_RIGHT or event.key == pg.K_d) and snake_dir != (-tile_size, 0):
                snake_dir = (tile_size, 0)

    """ Start Screen """
    if start_screen:
        game_mode = show_menu(screen).lower()
        if game_mode == 'quit':
            screen.fill('black')
            game_over = True
        elif game_mode == 'scoreboard':
            start_screen = False
            scoreboard = get_scoreboard()
            if scoreboard == []:
                scoreboard = [['No Scores Yet', '']]
                max_width = 0
            else:
                scoreboard.sort(key=lambda x: x[1], reverse=True)
                max_width = max([score_font.render(date, True, (255, 255, 255)).get_width() for date, score in scoreboard])
            screen.fill('black')
            x = 25
            y = 25
            for name, score in scoreboard:
                score_text = score_font.render(str(score), True, (255, 255, 255))
                name_text = score_font.render(name, True, (255, 255, 255))
                screen.blit(name_text, (x, y))
                screen.blit(score_text, (x + max_width + 20, y))
                y += 40

            pg.display.flip()

            while not start_screen:
                for event in pg.event.get():
                    if event.type == pg.KEYDOWN:
                        if event.key == pg.K_RETURN or event.key == pg.K_ESCAPE or event.key == pg.K_q:
                            start_screen = True
                            break
        else:
            snake_speed = mode[game_mode]['snake_speed']
            start_screen = False
            infinite = mode[game_mode]['infinite']
            snake.center = get_random_position()
            snake_dir = (0, 0)
            time = 0
            length = 1
            snake_body = [snake.copy()]
            score = 0
    else:
        screen.fill('black')
        pg.draw.rect(screen, food_color, food)
        [pg.draw.rect(screen, snake_color, body) for body in snake_body]

        """ Game Over Conditions """
        if snake.center in [body.center for body in snake_body[:-1]]:
            self_collision = True


        if snake.left < 0 or snake.right > window or snake.top < 0 or snake.bottom > window or self_collision:
            game_over_text = game_over_font.render('Game Over', True, game_over_color)
            game_over_text_rect = game_over_text.get_rect(center=(window/2, window/2))
            score_text = score_font.render(f'Score: {score}', True, score_color)
            score_text_rect = score_text.get_rect(center=(window/2, window/2 + 33))
            screen.fill('black')
            screen.blit(score_text, score_text_rect)
            screen.blit(game_over_text, game_over_text_rect)
            if event.type == pg.KEYDOWN:
                if event.key == pg.K_RETURN:
                    insert_scoreboard(score)
                    start_screen = True
                if event.key == pg.K_ESCAPE:
                    insert_scoreboard(score)
                    game_over = True
        if infinite:
            if snake.left < 0:
                snake.left = window - tile_size - 4
            if snake.right > window:
                snake.right = tile_size + 4
            if snake.top < 0:
                snake.top = window - tile_size - 4
            if snake.bottom > window:
                snake.bottom = tile_size + 4

        """ Game Logic """
        if snake.center == food.center:
            food.center = get_food_position(food)
            length += 1
            score += 1
            snake_speed = max(30, snake_speed - 3) if snake_speed > 60 else snake_speed - 1

        time_now = pg.time.get_ticks()

        if time_now - time > snake_speed:
            time = time_now
            snake.move_ip(snake_dir)
            snake_body.append(snake.copy())
            snake_body = snake_body[-length:]

        pg.display.flip()
        clock.tick(60)
